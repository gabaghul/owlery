package http

import (
	"net/http"

	"github.com/rs/zerolog"
)

type MailChimpAdapter struct {
	apiKey  string
	baseURL string
	server  string
	client  http.Client
	logger  *zerolog.Logger
}

type OmetriaEmailingIngestAdapter struct {
	apiKey  string
	baseURL string
	client  http.Client
	logger  *zerolog.Logger
}

func NewMailChimpAdapter(client http.Client, logger *zerolog.Logger, baseURL, apiKey, server string) MailChimpAdapter {
	return MailChimpAdapter{
		apiKey:  apiKey,
		baseURL: baseURL,
		server:  server,
		client:  client,
		logger:  logger,
	}
}

func NewOmetriaEmailingIngestAdapter(client http.Client, logger *zerolog.Logger, apiKey, baseUrl string) OmetriaEmailingIngestAdapter {
	return OmetriaEmailingIngestAdapter{
		client:  client,
		apiKey:  apiKey,
		baseURL: baseUrl,
		logger:  logger,
	}
}
