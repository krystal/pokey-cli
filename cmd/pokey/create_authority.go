package main

import (
	"errors"
	"fmt"

	"github.com/krystal/pokey-cli/internal/client"
	"github.com/krystal/pokey-cli/internal/configmanager"
	"github.com/krystal/pokey-cli/internal/pokey"
	"github.com/spf13/cobra"
)

var (
	createAuthorityLabel      string
	createAuthorityCommonName string
	createAuthorityYears      int
	createAuthorityKeySize    int
)

func runCreateAuthority(cmd *cobra.Command, args []string) error {
	apiHost := args[0]
	authorityName := args[1]

	if apiHost == "" {
		return errors.New("API host must be provided")
	}

	if authorityName == "" {
		return errors.New("Authority name must be provided")
	}

	authority, err := configManager.Authority(authorityName)
	if err != nil {
		return fmt.Errorf("could not read authority config: %w", err)
	}

	if authority != nil {
		return fmt.Errorf("an authority is already configured named %s", authorityName)
	}

	label := createAuthorityLabel
	if label == "" {
		label = createAuthorityCommonName
	}

	request := &client.CreateAuthorityRequest{
		Hostname: apiHost,
		Label:    label,
		KeySize:  createAuthorityKeySize,
		Years:    createAuthorityYears,
		Subject: &pokey.Subject{
			CommonName: createAuthorityCommonName,
		},
	}

	response := &client.CreateAuthorityResponse{}
	err = apiClient.MakeRequest(ctx, request, response)
	if err != nil {
		return fmt.Errorf("could not make the authority: %w", err)
	}

	fmt.Printf("Created new authority with ID %s\n", response.Authority.ID)
	fmt.Printf("Saved configuration as %s\n\n", authorityName)
	fmt.Printf("Get certificate using `pokey ca %s`\n", authorityName)

	authority = &configmanager.Authority{
		Hostname:    apiHost,
		ID:          response.Authority.ID,
		Certificate: response.Authority.Certificate,
		Secret:      response.Secret,
	}
	configManager.SaveAuthority(authorityName, authority)

	return nil
}

func init() {
	command := &cobra.Command{
		Use:   "create-authority <host> <name>",
		Short: "Create certificate authority",
		Args:  cobra.ExactArgs(2),
		RunE:  runCreateAuthority,
	}

	command.Flags().StringVarP(&createAuthorityLabel, "label", "l", "", "label for the CA (uses CN by default)")
	command.Flags().StringVar(&createAuthorityCommonName, "cn", "Certiticate Authority", "common name for CA")
	command.Flags().IntVar(&createAuthorityYears, "years", 30, "years for the CA")
	command.Flags().IntVar(&createAuthorityKeySize, "key-size", 4096, "size for the CA private key")

	rootCmd.AddCommand(command)
}
