package tlsutils

import (
	"fmt"
	"errors"
	"crypto/x509"
	"crypto/tls"
	"io/ioutil"
)


func LoadCACert(caFilename string) (*x509.CertPool, error) {
	caPem, err := ioutil.ReadFile(caFilename)

	if err != nil {
		return nil, errors.New(fmt.Sprintf(
			"tlsutils: Error reading cert file '%s': %v", caFilename, err))
	}

	certPool := x509.NewCertPool()

	if !certPool.AppendCertsFromPEM(caPem) {
		return nil, errors.New("tlsutils: Couldn't append cert from PEM")
	}
	return certPool, nil
}

func CreateTLSConfig(certPath, keyPath string) (*tls.Config, error) {
	// Load key and cert
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(
			"tlsutils: Error loading cert/key: '%s' '%s': %v", certPath,
			keyPath, err))
	}

	// Create TLS config for server
	return &tls.Config{ Certificates: []tls.Certificate{cert}}, nil
}
