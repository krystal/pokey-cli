package main

import (
	"context"
	"fmt"
	"os"

	"github.com/krystal/pokey-cli/internal/client"
	"github.com/krystal/pokey-cli/internal/configmanager"
	"github.com/spf13/cobra"
)

type contextKey int

const (
	ctxKey contextKey = iota
)

var (
	ctx           context.Context
	apiClient     *client.Client
	configManager *configmanager.Manager
	rootCmd       = &cobra.Command{
		Use:   "pokey",
		Short: "Pokey PKI tool",
		Long:  `Pokey is a command line tool for the pokey PKI API.`,
	}
)

func init() {
	cobra.OnInitialize(initConfigManager, initClient)
}

func initClient() {
	ctx = context.WithValue(context.Background(), ctxKey, "pokey")
	apiClient = client.New()
}

func initConfigManager() {
	configManager = configmanager.New(
		fmt.Sprintf("%s/.pokey", os.Getenv("HOME")),
	)
}

func main() {
	rootCmd.Execute()
}
