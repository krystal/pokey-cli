package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

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
	configRoot    string
	rootCmd       = &cobra.Command{
		Use:          "pokey",
		Short:        "Pokey PKI tool",
		Long:         `Pokey is a command line tool for the pokey PKI API.`,
		SilenceUsage: true,
	}
)

func init() {
	cobra.OnInitialize(initConfigManager, initClient)

	rootFlags := rootCmd.PersistentFlags()
	rootFlags.StringVar(&configRoot, "config-root", fmt.Sprintf("%s/.pokey", os.Getenv("HOME")),
		"Path to configuration files")
}

func initClient() {
	ctx = context.WithValue(context.Background(), ctxKey, "pokey")
	apiClient = client.New()
}

func initConfigManager() {
	path, _ := filepath.Abs(configRoot)
	configManager = configmanager.New(path)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
