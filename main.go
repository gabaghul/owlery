package main

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	httpAdapter "github.com/gabaghul/owlery/src/adapters/http"
	psqlAdapter "github.com/gabaghul/owlery/src/adapters/psql"
	redisAdapter "github.com/gabaghul/owlery/src/adapters/redis"
	"github.com/gabaghul/owlery/src/domain/emailing/services"
	"github.com/gabaghul/owlery/src/helpers"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	ctx := context.Background()
	configs := helpers.LoadConfigs(fmt.Sprintf("%s/configs", basepath), "application-local")
	logger := helpers.GetLogger()
	httpClient := helpers.NewHTTPClient()

	psql, err := psqlAdapter.NewPsqlAdapter(&logger,
		configs.Postgres.Host,
		configs.Postgres.Port,
		configs.Postgres.Username,
		configs.Postgres.Password,
		configs.Postgres.Database,
	)
	if err != nil {
		panic(fmt.Sprintf("cannot start psql adapter: %s", err))
	}

	redis := redisAdapter.NewRedisAdapter(&logger,
		configs.Redis.Host,
		configs.Redis.Port,
		configs.Redis.Password,
		configs.Redis.Database,
	)

	mailchimp := httpAdapter.NewMailChimpAdapter(httpClient,
		&logger,
		configs.HTTP.Mailchimp.BaseURL,
		configs.HTTP.Mailchimp.APIKey,
		configs.HTTP.Mailchimp.Server,
	)
	ometria := httpAdapter.NewOmetriaAdapter(httpClient,
		&logger,
		configs.HTTP.Ometria.APIKey,
		configs.HTTP.Ometria.BaseURL,
	)

	service := services.NewEmailingService(&logger, psql, redis, mailchimp, ometria)

	if err = service.DoEmailPooling(ctx); err != nil {
		logger.Err(err).Msg("error processing emailing pooling")
	}
}
