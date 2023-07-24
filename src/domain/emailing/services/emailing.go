package services

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/gabaghul/owlery/src/domain/emailing/models"
	"github.com/pkg/errors"
)

const defaultCount = 500 // defines the maximum amount of emails fetched from mailchimp per request

func (s EmailingService) DoEmailPooling(ctx context.Context) error {
	configs, err := s.psql.GetAllEmailingConfigs(ctx)
	if err != nil {
		return errors.Wrap(err, "could not fetch emailing configs from sql adapter")
	}

	channelResults := make(chan string)
	var wg sync.WaitGroup

	for _, config := range configs {
		// adding one process to wait group counter
		wg.Add(1)

		go s.processPooling(ctx, config, &wg, channelResults) // start go routines for each config
	}

	// closes channels after all processing is done
	go func() {
		wg.Wait() // wait for all go routines to finish
		close(channelResults)
	}()

	return nil
}

func (s EmailingService) processPooling(ctx context.Context, config models.EmailingConfigs, wg *sync.WaitGroup, channelResults chan<- string) {
	defer wg.Done()
	now := time.Now()

	contactLists, err := s.psql.GetContactListsByClientID(ctx, config.ClientID)
	if err != nil {
		s.logger.Err(err).
			Str("client_id", strconv.Itoa(int(config.ClientID))).
			Str("at", time.Now().Format(time.RFC3339)).
			Msg("could not get contacts list from sql adapter")

		channelResults <- "error"
		return
	}

	for _, list := range contactLists {
		offset, err := s.redis.GetEmailingMemberListOffset(ctx, list.ClientID, list.ListID)
		if err != nil {
			s.logger.Err(err).
				Str("client_id", strconv.Itoa(int(config.ClientID))).
				Str("list_id", list.ListID).
				Str("at", time.Now().Format(time.RFC3339)).
				Msg("could not get offset for member list api pagination")

			continue
		}

		contacts, err := s.mailchimp.GetContactsByListID(ctx, list.ListID, int64(offset), int64(15))
		if err != nil {
			s.logger.Err(err).
				Str("client_id", strconv.Itoa(int(config.ClientID))).
				Str("list_id", list.ListID).
				Str("at", time.Now().Format(time.RFC3339)).
				Msg("could not get contacts by list id from members list api")

			continue
		}

		fmt.Println("WORKED!!", contacts)
	}

	s.logger.Debug().
		Str("client_id", strconv.Itoa(int(config.ClientID))).
		Str("elapsed_time", fmt.Sprintf("%dms", time.Now().UnixMilli()-now.UnixMilli())).
		Msg("successfully processed emailing config")
	channelResults <- "done" // sending random info to the channel just to have some output
}
