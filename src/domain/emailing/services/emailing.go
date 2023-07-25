package services

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/gabaghul/owlery/src/adapters/http"
	"github.com/gabaghul/owlery/src/adapters/redis"
	"github.com/gabaghul/owlery/src/domain/emailing/models"
	"github.com/pkg/errors"
)

const (
	defaultCount    = 500 // defines the maximum amount of emails fetched from mailchimp per request
	firstNameMapKey = "FNAME"
	lastNameMapKey  = "LNAME"
)

func (s EmailingService) DoEmailPooling(ctx context.Context) error {
	configs, err := s.psql.GetAllEmailingConfigs(ctx)
	if err != nil {
		return errors.Wrap(err, "could not fetch emailing configs from sql adapter")
	}

	var wg sync.WaitGroup

	for _, config := range configs {
		// adding one process to wait group counter
		wg.Add(1)

		go s.processPooling(ctx, config, &wg) // start go routines for each config
	}

	wg.Wait()
	return nil
}

func (s EmailingService) processPooling(ctx context.Context, config models.EmailingConfig, wg *sync.WaitGroup) {
	defer wg.Done()
	now := time.Now()

	contactLists, err := s.psql.GetContactListsByClientID(ctx, config.ClientID)
	if err != nil {
		s.logger.Err(err).
			Int64("client_id", config.ClientID).
			Str("at", time.Now().Format(time.RFC3339)).
			Msg("could not get contacts list from sql adapter")

		return
	}

	for _, list := range contactLists {
		offset, err := s.redis.GetEmailingMemberListOffset(ctx, list.ClientID, list.ListID)
		if err != nil {
			s.logger.Err(err).
				Int64("client_id", config.ClientID).
				Str("list_id", list.ListID).
				Str("at", time.Now().Format(time.RFC3339)).
				Msg("could not get offset for member list api pagination")

			continue
		}

		for {
			var contacts []models.Contact
			contactsBatch, err := s.mailchimp.GetContactsByListID(ctx, list.ListID, int64(offset), defaultCount)
			if err != nil {
				s.logger.Err(err).
					Int64("client_id", config.ClientID).
					Str("list_id", list.ListID).
					Str("at", time.Now().Format(time.RFC3339)).
					Msg("could not get contacts by list id from members list api")

				break
			}
			contacts = append(contacts, s.toDomain(contactsBatch, list.ClientID)...)

			_, err = s.ometria.IngestContactRecords(ctx, contacts)
			if err != nil {
				s.logger.Err(err).
					Int64("client_id", config.ClientID).
					Str("list_id", list.ListID).
					Str("at", time.Now().Format(time.RFC3339)).
					Msg("could not send contacts list to ometria api")

				break
			}

			offset += defaultCount
			if offset >= int(contactsBatch.TotalItems) {
				s.logger.Info().
					Int64("client_id", config.ClientID).
					Str("total_processed", strconv.Itoa(int(contactsBatch.TotalItems))).
					Msg("fully processed contacts lists for client id")

				offset = int(contactsBatch.TotalItems)
				break
			}
		}
		if err != nil {
			s.logger.Err(err).
				Int64("client_id", config.ClientID).
				Str("list_id", list.ListID).
				Str("at", time.Now().Format(time.RFC3339)).
				Msg("could not ingest data for ometria")

			continue
		}

		err = s.redis.Store(ctx, redis.EmailingMemberListOffsetKey(list.ClientID, list.ListID), offset, 0)
		if err != nil {
			s.logger.Err(err).
				Int64("client_id", config.ClientID).
				Str("list_id", list.ListID).
				Int("offset", offset).
				Str("at", time.Now().Format(time.RFC3339)).
				Msg("could not update offset for next scheduled task")
		}
	}

	s.logger.Debug().
		Str("client_id", strconv.Itoa(int(config.ClientID))).
		Str("elapsed_time", fmt.Sprintf("%dms", time.Now().UnixMilli()-now.UnixMilli())).
		Msg("successfully processed emailing config")
}

func (s EmailingService) toDomain(contact http.GetContactsByListIDResponse, clientID int64) []models.Contact {
	contacts := make([]models.Contact, len(contact.Members))
	for i, c := range contact.Members {
		var firstName, lastName string
		if v, ok := c.MergeFields[firstNameMapKey]; ok {
			firstName = v.(string)
		}
		if v, ok := c.MergeFields[lastNameMapKey]; ok {
			lastName = v.(string)
		}
		contacts[i] = models.Contact{
			ID:        c.UniqueEmailID,
			ClientID:  clientID,
			Firstname: firstName,
			Lastname:  lastName,
		}
	}

	return contacts
}
