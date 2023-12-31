package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gabaghul/owlery/src/domain/emailing/models"
	"github.com/pkg/errors"
)

func (a PsqlAdapter) GetAllEmailingConfigs(ctx context.Context) (configs []models.EmailingConfig, err error) {
	rows, err := a.pool.QueryContext(ctx, "SELECT client_id, created_at, updated_at, active FROM emailing_configs")
	if err != nil && err != sql.ErrNoRows {
		return []models.EmailingConfig{}, errors.Wrap(err, "could not fetch all emailing configs")
	}
	defer rows.Close()

	for rows.Next() {
		config := models.EmailingConfig{}
		if err := rows.Scan(
			&config.ClientID,
			&config.CreatedAt,
			&config.UpdatedAt,
			&config.Active,
		); err != nil {
			return []models.EmailingConfig{}, errors.Wrap(err, "could not scan all returned values from emailing configs table")
		}

		configs = append(configs, config)
	}

	return configs, nil
}

func (a PsqlAdapter) GetEmailingConfigsByClientID(ctx context.Context, clientID int64) (configs []models.EmailingConfig, err error) {
	rows, err := a.pool.QueryContext(ctx, "SELECT client_id, created_at, updated_at, active FROM emailing_configs WHERE client_id = $1", clientID)
	if err != nil && err != sql.ErrNoRows {
		return []models.EmailingConfig{}, errors.Wrap(err, fmt.Sprintf("could not fetch emailing configs for client id %d", clientID))
	}
	defer rows.Close()

	for rows.Next() {
		config := models.EmailingConfig{}
		if err := rows.Scan(
			&config.ClientID,
			&config.CreatedAt,
			&config.UpdatedAt,
			&config.Active,
		); err != nil {
			return []models.EmailingConfig{}, errors.Wrap(err, fmt.Sprintf("could not scan filtered returned values from emailing configs table for client id %d", clientID))
		}

		configs = append(configs, config)
	}

	return configs, nil
}
