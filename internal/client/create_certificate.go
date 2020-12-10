package client

import (
	"context"
	"fmt"

	"github.com/krystal/pokey-cli/internal/pokey"
)

type CreateCertificateRequest struct {
	AuthorityID string         `json:"-"`
	Subject     *pokey.Subject `json:"subject,omitempty"`
	Hostname    string         `json:"hostname,omitempty"`
	Secret      string         `json:"secret,omitempty"`
	Usage       string         `json:"usage,omitempty"`
	SANs        *SANsHash      `json:"sans,omitempty"`
}

type SANsHash struct {
	DNS []string `json:"dns"`
	IP  []string `json:"ip"`
}

func (s *CreateCertificateRequest) URL() string {
	return fmt.Sprintf("https://%s/v1/authorities/%s/certificates", s.Hostname, s.AuthorityID)
}

func (s *CreateCertificateRequest) HTTPMethod() string {
	return "POST"
}

type CreateCertificateResponse struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

func (s *Client) CreateCertificate(ctx context.Context, request *CreateCertificateRequest) (*CreateCertificateResponse, error) {
	response := &CreateCertificateResponse{}
	err := s.MakeRequest(ctx, request, response)
	if err != nil {
		return response, fmt.Errorf("could not make the authority: %w", err)
	}

	return response, nil
}
