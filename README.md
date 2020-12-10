# Pokey CLI

This is a CLI for talking to a Pokey PKI API.

## Usage

```bash
# Begin by creating an authority by providing the host for the PKI
# service and the name you wish to use to identify the authroity.
$ pokey create-authority pki.infra.katapult.io example-ca
# => Authority created with ID abcdef1234
# => Configuration stored in ~/.pokey/authorities/example-ca

# Get a certificate authority file
$ pokey get-authority-cert example-ca
# => -----BEGIN CERTIFICATE----...

# Delete a certificate authority
$ pokey delete-authority example-ca
# => Authority abcdef1234 deleted, configuration removed

# Create a certificate
$ pokey create-certificate example-ca example.com \
              --export-to path/to/export/demo-cert
              --usage web-server
              --ip-sans 10.2.1.1,2.3.1.2
              --dns-sans www.example.com,mail.example.com
              --country GB
              --state Wiltshire
              --locality Salisbury
              --organization Krystal
              --organization-unit 'Software Engineering'
# => Certificate written to path/to/export/demo-cert.cert.pem
# => Private key written to path/to/export/demo-cert.key.pem

# Get a list of certificates
$ pokey list-certificates example-ca
# => 0001    /CN=example.com      Exp: 2020-02-02 12:33
# => 0002    /CN=10.3.1.2         Exp: 2020-02-02 12:56
```