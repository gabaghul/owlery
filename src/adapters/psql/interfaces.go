package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type PsqlAdapter struct {
	pool   *sql.DB
	logger *zerolog.Logger
}

func NewPsqlAdapter(logger *zerolog.Logger, host string, port int, user, password, dbname string) (PsqlAdapter, error) {
	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	pool, err := sql.Open("postgres", psqlConnString)
	if err != nil {
		return PsqlAdapter{}, errors.Wrap(err, "could not connect to database")
	}

	return PsqlAdapter{
		pool:   pool,
		logger: logger,
	}, nil
}
