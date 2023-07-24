package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

const (
	defaultOffset = int64(0)
	defaultCount  = int64(1000)
	fieldsToFetch = "total_items,members.unique_email_id,members.merge_fields,members.email_address,members.status"
)

type GetContactsByListIDResponse struct {
	Members    []GetContactsByListIDMembers `json:"members"`
	TotalItems int64                        `json:"total_items"`
}

type GetContactsByListIDErrorResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

type GetContactsByListIDMembers struct {
	UniqueEmailID string                 `json:"unique_email_id"`
	MergeFields   map[string]interface{} `json:"merge_fields"`
	Email         string                 `json:"email_address"`
	Status        string                 `json:"status"`
}

func (a MailChimpAdapter) GetContactsByListID(ctx context.Context, listID string, offset, count int64) (response GetContactsByListIDResponse, err error) {
	if offset <= 0 {
		offset = defaultOffset
	}
	if count <= 0 {
		count = 1000
	}

	url := fmt.Sprintf("%s/lists/%s/members?offset=%d&count=%d&fields=%s", a.BaseURL, listID, offset, count, fieldsToFetch)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return GetContactsByListIDResponse{}, errors.Wrap(err, "could not create request for get contacts by list id resource")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.APIKey))
	req.Header.Add("server", a.Server)

	resBody, err := a.Client.Do(req)
	if err != nil {
		return GetContactsByListIDResponse{}, errors.Wrap(err, "could not receive data from get contacts by list id resource")
	}
	defer resBody.Body.Close()

	switch resBody.StatusCode {
	case http.StatusOK:
		body, err := ioutil.ReadAll(resBody.Body)
		if err != nil {
			return GetContactsByListIDResponse{}, errors.Wrap(err, "could not retrieve response body content from members list info callback")
		}
		if err := json.Unmarshal(body, &response); err != nil {
			return GetContactsByListIDResponse{}, errors.Wrap(err, "could not unmarshal response body from members list info callback to json")
		}

		a.Logger.Debug().
			Str("list_id", listID).
			Str("total_items", strconv.Itoa(int(response.TotalItems))).
			Msg("successfully received data from members list info api")
		return response, nil
	default:
		body, err := ioutil.ReadAll(resBody.Body)
		if err != nil {
			return GetContactsByListIDResponse{}, errors.Wrap(err, "could not retrieve error response body content from members list info callback")
		}

		var errorResponse GetContactsByListIDErrorResponse
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return GetContactsByListIDResponse{}, errors.Wrap(err, "could not unmarshal error response body from members list info callback to json")
		}

		a.Logger.Error().
			Str("list_id", listID).
			Str("type", errorResponse.Type).
			Str("title", errorResponse.Title).
			Str("status", strconv.Itoa(errorResponse.Status)).
			Str("detail", errorResponse.Detail).
			Str("instance", errorResponse.Instance).
			Msg("error received from members list info")
		return GetContactsByListIDResponse{}, errors.New(fmt.Sprintf("error with status code %d received from members list info api", resBody.StatusCode))
	}
}
