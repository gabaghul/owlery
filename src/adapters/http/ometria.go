package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gabaghul/owlery/src/domain/emailing/models"
	"github.com/pkg/errors"
)

type ingestContactRecordsEntity struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Status    string `json:"subscribed"`
}

type ingestContactRecordsResponse struct {
	Status   string `json:"status"`
	Response int    `json:"response"`
}

type ingestContactRecordsErrorResponse struct {
	Status int    `json:"status"`
	Reason string `json:"reason"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (a OmetriaAdapter) IngestContactRecords(ctx context.Context, contacts []models.Contact) (int, error) {
	url := fmt.Sprintf("%s/record", a.baseURL)

	data := a.toRequestBody(contacts)
	reqBody, err := json.Marshal(data)
	if err != nil {
		return 0, errors.Wrap(err, "could not marshal request to json for ingest contacts api")
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return 0, errors.Wrap(err, "could not create request for ingest contacts api")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", a.apiKey)

	resBody, err := a.client.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "error calling ingest contacts api")
	}
	defer resBody.Body.Close()
	switch resBody.StatusCode {
	case http.StatusCreated:
		var response ingestContactRecordsResponse

		body, err := ioutil.ReadAll(resBody.Body)
		if err != nil {
			return 0, errors.Wrap(err, "could not retrieve response body content from contact ingestion callback")
		}
		if err := json.Unmarshal(body, &response); err != nil {
			return 0, errors.Wrap(err, "could not unmarshal response body from contact ingestion callback to json")
		}

		a.logger.Debug().
			Str("total_items", strconv.Itoa(int(response.Response))).
			Msg("successfully posted data to contact records api")
		return response.Response, nil
	default:
		var response ingestContactRecordsErrorResponse

		body, err := ioutil.ReadAll(resBody.Body)
		if err != nil {
			return 0, errors.Wrap(err, "could not retrieve response error body content from contact ingestion callback")
		}
		if err := json.Unmarshal(body, &response); err != nil {
			return 0, errors.Wrap(err, "could not unmarshal response error body from contact ingestion callback to json")
		}

		a.logger.Error().
			Str("status", strconv.Itoa(int(response.Status))).
			Str("reason", response.Reason).
			Str("title", response.Title).
			Str("detail", response.Detail).
			Msg(fmt.Sprintf("error with status code %d received from contact ingest api", resBody.StatusCode))
		return 0, errors.New(fmt.Sprintf("error with status code %d received from contact ingest api"))
	}
}

func (a OmetriaAdapter) toRequestBody(contacts []models.Contact) []ingestContactRecordsEntity {
	entities := make([]ingestContactRecordsEntity, len(contacts))
	for i, contact := range contacts {
		entities[i] = ingestContactRecordsEntity{
			ID:        contact.ID,
			Firstname: contact.Firstname,
			Lastname:  contact.Lastname,
			Email:     contact.Email,
			Status:    contact.Status,
		}
	}

	return entities
}
