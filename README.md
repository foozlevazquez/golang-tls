# golang-tls
Examples of TLS usage in golang


## Setup
### Generating test pki

From `golang-tls` directory.
```
  git clone git@github.com:OpenVPN/easy-rsa.git test_pki

  cd test_pki/easyrsa3
  ./easyrsa init-pki

  # Note CA password used here
  ./easyrsa build-ca

  # Use CA password here (server cert has no password)
  ./easyrsa build-server-full server nopass

  # Use CA password here (client cert has no password)
  ./easyrsa build-client-full client nopass

  # Check certs are valid against CA:
  openssl verify -CAfile ./pki/ca.crt ./pki/issued/*.crt
```
