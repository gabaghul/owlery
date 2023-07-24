package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gabaghul/owlery/src/domain/emailing/models"
	"github.com/pkg/errors"
)

func (a PsqlAdapter) GetAllContactLists(ctx context.Context) (lists []models.ContactList, err error) {
	rows, err := a.pool.QueryContext(ctx, "SELECT list_id, client_id, created_at, updated_at, active FROM contact_lists")
	if err != nil && err != sql.ErrNoRows {
		return []models.ContactList{}, errors.Wrap(err, "could not fetch all contact lists")
	}
	defer rows.Close()

	for rows.Next() {
		list := models.ContactList{}
		if err := rows.Scan(
			&list.ListID,
			&list.ClientID,
			&list.CreatedAt,
			&list.UpdatedAt,
			&list.Active,
		); err != nil {
			return []models.ContactList{}, errors.Wrap(err, "could not scan all returned values from contact lists table")
		}

		lists = append(lists, list)
	}

	return lists, nil
}

func (a PsqlAdapter) GetContactListsByClientID(ctx context.Context, clientID int64) (lists []models.ContactList, err error) {
	rows, err := a.pool.QueryContext(ctx, "SELECT list_id, client_id, created_at, updated_at, active FROM contact_lists WHERE client_id = $1", clientID)
	if err != nil && err != sql.ErrNoRows {
		return []models.ContactList{}, errors.Wrap(err, fmt.Sprintf("could not fetch contact lists for client id %d", clientID))
	}
	defer rows.Close()

	for rows.Next() {
		list := models.ContactList{}
		if err := rows.Scan(
			&list.ListID,
			&list.ClientID,
			&list.CreatedAt,
			&list.UpdatedAt,
			&list.Active,
		); err != nil {
			return []models.ContactList{}, errors.Wrap(err, fmt.Sprintf("could not scan filtered returned values from contact lists table for client id %d", clientID))
		}

		lists = append(lists, list)
	}

	return lists, nil
}
