package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	hostname   string
	httpClient *http.Client
}

func (s *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	r, err := s.httpClient.Do(req)
	if err != nil {
		return r, err
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return r, err
	}

	if r.StatusCode < 200 || r.StatusCode >= 300 {
		return r, errors.New("Not 2xx")
	}

	err = json.Unmarshal(body, v)

	return r, err
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 60,
		},
	}
}

func (s *Client) MakeRequest(ctx context.Context, request Request, response interface{}) error {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		request.HTTPMethod(),
		request.URL(),
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	_, err = s.Do(req, response)

	return err
}

type Request interface {
	URL() string
	HTTPMethod() string
}
