package http

import (
	"net/http"

	"github.com/rs/zerolog"
)

type MailChimpAdapter struct {
	APIKey  string
	BaseURL string
	Server  string
	Client  http.Client
	Logger  *zerolog.Logger
}

func NewMailChimpAdapter(client http.Client, logger *zerolog.Logger, baseURL, apiKey, server string) MailChimpAdapter {
	return MailChimpAdapter{
		APIKey:  apiKey,
		BaseURL: baseURL,
		Server:  server,
		Client:  client,
		Logger:  logger,
	}
}
