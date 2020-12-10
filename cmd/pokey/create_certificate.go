package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/krystal/pokey-cli/internal/client"
	"github.com/krystal/pokey-cli/internal/pokey"
	"github.com/spf13/cobra"
)

var (
	createCertificateExportPathFlag       string
	createCertificateCountryFlag          string
	createCertificateStateFlag            string
	createCertificateLocalityFlag         string
	createCertificateOrganizationFlag     string
	createCertificateOrganizationUnitFlag string
	createCertificateUsageFlag            string
	createCertificateDNSSANsFlag          string
	createCertificateIPSANsFlag           string
)

func runCreateCertficate(cmd *cobra.Command, args []string) error {
	authorityName := args[0]
	commonName := args[1]

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

	if commonName == "" {
		return errors.New("common name must be provided")
	}

	dnsSANs := createCertificateDNSSANsFlag
	var dnsSANsAsArray []string
	if dnsSANs != "" {
		dnsSANsAsArray = strings.Split(dnsSANs, ",")
	}

	ipSANs := createCertificateIPSANsFlag
	var ipSANsAsArray []string
	if ipSANs != "" {
		ipSANsAsArray = strings.Split(ipSANs, ",")
	}

	request := &client.CreateCertificateRequest{
		AuthorityID: authority.ID,
		Hostname:    authority.Hostname,
		Secret:      authority.Secret,
		Usage:       strings.ReplaceAll(createCertificateUsageFlag, "-", "_"),
		SANs: &client.SANsHash{
			DNS: dnsSANsAsArray,
			IP:  ipSANsAsArray,
		},
		Subject: &pokey.Subject{
			CommonName:       commonName,
			Country:          createCertificateCountryFlag,
			State:            createCertificateStateFlag,
			Locality:         createCertificateLocalityFlag,
			Organization:     createCertificateOrganizationFlag,
			OrganizationUnit: createCertificateOrganizationUnitFlag,
		},
	}
	response, err := apiClient.CreateCertificate(ctx, request)
	if err != nil {
		return fmt.Errorf("could not create certificate: %w", err)
	}

	if createCertificateExportPathFlag == "" {
		fmt.Print(response.PrivateKey)
		fmt.Print(response.Certificate)

		return nil
	}

	certificatePath := fmt.Sprintf("%s.cert.pem", createCertificateExportPathFlag)
	privateKeyPath := fmt.Sprintf("%s.key.pem", createCertificateExportPathFlag)

	err = ioutil.WriteFile(certificatePath, []byte(response.Certificate), 0644) //nolint:gosec
	if err != nil {
		return fmt.Errorf("error writing cert file: %w", err)
	}

	err = ioutil.WriteFile(privateKeyPath, []byte(response.PrivateKey), 0600)
	if err != nil {
		return fmt.Errorf("error writing private key file: %w", err)
	}

	fmt.Printf("Written certificate to %s\n", certificatePath)
	fmt.Printf("Written private key to %s\n", privateKeyPath)

	return nil
}

func init() {
	command := &cobra.Command{
		Use:   "cert <ca> <common-name>",
		Short: "Create certificate",
		Args:  cobra.ExactArgs(2),
		RunE:  runCreateCertficate,
	}

	command.Flags().StringVarP(&createCertificateExportPathFlag, "export-path", "e", "", "path to export to (defaults to STDOUT)")
	command.Flags().StringVarP(&createCertificateCountryFlag, "country", "c", "", "country code")
	command.Flags().StringVarP(&createCertificateStateFlag, "state", "s", "", "state")
	command.Flags().StringVarP(&createCertificateLocalityFlag, "locality", "l", "", "locality")
	command.Flags().StringVarP(&createCertificateOrganizationFlag, "org", "o", "", "organization")
	command.Flags().StringVarP(&createCertificateOrganizationUnitFlag, "org-unit", "t", "", "organization unit")
	command.Flags().StringVarP(&createCertificateUsageFlag, "usage", "u", "", "usage")
	command.Flags().StringVarP(&createCertificateDNSSANsFlag, "dns-sans", "", "", "DNS SANs")
	command.Flags().StringVarP(&createCertificateIPSANsFlag, "ip-sans", "", "", "IP SANs")

	rootCmd.AddCommand(command)
}
