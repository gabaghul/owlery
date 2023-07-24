package psql

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type PsqlAdapter struct {
	Pool   *sql.DB
	Logger *zerolog.Logger
}

func NewPsqlAdapter(logger *zerolog.Logger, host, port, user, password, dbname string) (PsqlAdapter, error) {
	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	pool, err := sql.Open("postgres", psqlConnString)
	if err != nil {
		return PsqlAdapter{}, errors.Wrap(err, "could not connect to database")
	}

	return PsqlAdapter{
		Pool:   pool,
		Logger: logger,
	}, nil
}
