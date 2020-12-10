package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/krystal/pokey-cli/internal/client"
	"github.com/spf13/cobra"
)

func runDeleteAuthority(cmd *cobra.Command, args []string) error {
	authorityName := args[0]

	if authorityName == "" {
		return errors.New("authority name must be provided")
	}

	authority, err := configManager.Authority(authorityName)
	if err != nil {
		return fmt.Errorf("could not read authority config: %w", err)
	}

	if authority == nil {
		return fmt.Errorf("no authority exists with name %s", authorityName)
	}

	request := &client.DeleteAuthorityRequest{
		ID:       authority.ID,
		Hostname: authority.Hostname,
		Secret:   authority.Secret,
	}

	_, err = apiClient.DeleteAuthority(ctx, request)
	if err != nil {
		return fmt.Errorf("could not delete on api: %w", err)
	}

	err = os.Remove(authority.ConfigPath)
	if err != nil {
		return fmt.Errorf("could not delete config at path %s: %w", authority.ConfigPath, err)
	}

	fmt.Printf("Deleted authority with ID %s\n", authority.ID)
	fmt.Printf("Deleted config from %s\n", authority.ConfigPath)

	return nil
}

func init() {
	command := &cobra.Command{
		Use:   "delete-authority <name>",
		Short: "Delete certificate authority",
		Args:  cobra.ExactArgs(1),
		RunE:  runDeleteAuthority,
	}

	rootCmd.AddCommand(command)
}
