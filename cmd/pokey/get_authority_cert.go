package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var getAuthorityCertExplain bool

func runGetAuthorityCert(cmd *cobra.Command, args []string) error {
	authorityName := args[0]

	if authorityName == "" {
		return errors.New("authority name must be provided")
	}

	authority, err := configManager.Authority(authorityName)
	if err != nil {
		return fmt.Errorf("could not read authority config: %w", err)
	}

	if authority == nil {
		return fmt.Errorf("not authority exists with name %s", authorityName)
	}

	if getAuthorityCertExplain {
		cmd := exec.Command("openssl", "x509", "-noout", "-text")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}

		defer stdin.Close()
		io.WriteString(stdin, authority.Certificate)

		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(out))
	} else {
		fmt.Println(strings.TrimSpace(authority.Certificate))
	}

	return nil
}

func init() {
	command := &cobra.Command{
		Use:   "ca <name>",
		Short: "Get certificate authority certificate",
		Args:  cobra.ExactArgs(1),
		RunE:  runGetAuthorityCert,
	}

	command.Flags().BoolVarP(&getAuthorityCertExplain, "explain", "e", false, "explain the certificate")

	rootCmd.AddCommand(command)
}
