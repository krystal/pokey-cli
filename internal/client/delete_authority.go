package client

import (
	"context"
	"fmt"
)

type DeleteAuthorityResponse struct {
	Status bool
}

type DeleteAuthorityRequest struct {
	Hostname string `json:"-"`
	Secret   string `json:"secret,omitempty"`
	ID       string `json:"-"`
}

func (s *DeleteAuthorityRequest) URL() string {
	return fmt.Sprintf("https://%s/v1/authorities/%s", s.Hostname, s.ID)
}

func (s *DeleteAuthorityRequest) HTTPMethod() string {
	return "DELETE"
}

func (s *Client) DeleteAuthority(ctx context.Context, request *DeleteAuthorityRequest) (*DeleteAuthorityResponse, error) {
	response := &DeleteAuthorityResponse{}
	err := s.MakeRequest(ctx, request, response)
	if err != nil {
		return response, fmt.Errorf("could not delete the authority: %w", err)
	}

	return response, nil
}
