package services

import (
	"context"
	"time"

	"github.com/gabaghul/owlery/src/adapters/http"
	"github.com/gabaghul/owlery/src/domain/emailing/models"
	"github.com/rs/zerolog"
)

type RedisAdapter interface {
	Store(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	GetEmailingMemberListOffset(ctx context.Context, clientID int64, listID string) (int, error)
}

type MailChimpAdapter interface {
	GetContactsByListID(ctx context.Context, listID string, offset, count int64) (http.GetContactsByListIDResponse, error)
}

type PsqlAdapter interface {
	GetAllContactLists(ctx context.Context) ([]models.ContactList, error)
	GetContactListsByClientID(ctx context.Context, clientID int64) ([]models.ContactList, error)
	GetAllEmailingConfigs(ctx context.Context) ([]models.EmailingConfig, error)
	GetEmailingConfigsByClientID(ctx context.Context, clientID int64) ([]models.EmailingConfig, error)
}

type EmailingService struct {
	logger    *zerolog.Logger
	psql      PsqlAdapter
	redis     RedisAdapter
	mailchimp MailChimpAdapter
}

func NewEmailingService(logger *zerolog.Logger, psql PsqlAdapter, redis RedisAdapter, mailchimp MailChimpAdapter) EmailingService {
	return EmailingService{
		logger:    logger,
		psql:      psql,
		redis:     redis,
		mailchimp: mailchimp,
	}
}
