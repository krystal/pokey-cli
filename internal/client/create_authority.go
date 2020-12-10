package client

import (
	"context"
	"fmt"

	"github.com/krystal/pokey-cli/internal/pokey"
)

type CreateAuthorityResponse struct {
	Authority *pokey.Authority `json:"authority,omitempty"`
	Secret    string           `json:"secret,omitempty"`
}

type CreateAuthorityRequest struct {
	Hostname string         `json:"-"`
	Subject  *pokey.Subject `json:"subject,omitempty"`
	Label    string         `json:"label,omitempty"`
	Years    int            `json:"years,omitempty"`
	KeySize  int            `json:"key_size,omitempty"`
}

func (s *CreateAuthorityRequest) URL() string {
	return fmt.Sprintf("https://%s/v1/authorities", s.Hostname)
}

func (s *CreateAuthorityRequest) HTTPMethod() string {
	return "POST"
}

func (s *Client) CreateAuthority(ctx context.Context, request *CreateAuthorityRequest) (*CreateAuthorityResponse, error) {
	response := &CreateAuthorityResponse{}
	err := s.MakeRequest(ctx, request, response)
	if err != nil {
		return response, fmt.Errorf("could not make the authority: %w", err)
	}

	return response, nil
}
