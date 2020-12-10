# Pokey CLI

This is a CLI for talking to a Pokey PKI API.

Remember, this is a very simple CA system designed for a very specific use case. There is no certificate revokation, no CSRs and no funny business. You just make a CA and ask it to give you certificates. The CA will immediately respond with signed certificate and a private key for your use.

## Install

1. Download the latest binary from the GitHub releases page
2. Extract it
3. Copy `pokey` into `/usr/local/bin`
4. Set permissions with `chmod +x /usr/local/bin`
5. Run it

## Usage

1. Begin by creating a certificate authority

  ```
  $ pokey create-authority pki.yourdomain.com example-ca
  ```

  You should replace the `example-ca` with a tag that you want to use to reference this authority when issuing certificates.

  You can pass additional options such as `--cn` to specify a custom common name for the CA certificate. You can also use `--years` and `--key-size` - by default CAs are valid for 30 days with a 4096-bit private key. The CA private key will never leave the PKI server.

2. Next, you can get a copy of your CA file.

  ```
  $ pokey ca example-ca
  ```

3. Create a certificate. At the most basic, you can just provide the name of your CA and the common name for the certificate. This will generate the cert and print both the private key and certificate to your STDOUT. **Note, the private key that is generated is not stored on the server and cannot be retreived if you lose it.**

  ```
  $ pokey cert example-ca example.com
  ```

  You can choose to export the key & certificate to your file system. Using this option, two files will be created with a suffix of `.key.pem` and `.cert.pem` as appropriate.

  ```
  $ pokey cert example-ca example.com -e path/to/file
  ```

  You can pass additional options to add more information (such as country, state, organization etc...). See the `--help` output for details.

  If you wish to add SANs to your certificate, you can specify these:

  ```
  $ pokey cert example-ca example.com --dns-sans www.example.com,mail.example.com
  $ pokey cert example-ca 10.0.0.1 --ip-sans 10.0.0.2,10.0.0.3
  ```

  You can also set the usage for the key depending on whether you're planning to use this for server or client authentication.

  ```
  $ pokey cert example-ca example.com -u web-server
  $ pokey cert example-ca example.com -u web-client
  ```
