package main

import (
	"time"
	"net"

	"github.com/foozlevazquez/golang-tls/tlsutils"
)


const (
	CACertPath = "./test_pki/easyrsa3/pki/ca.crt"

	ClientCertPath = "./test_pki/easyrsa3/pki/client.crt"
	ClientKeyPath  = "./test_pki/easyrsa3/pki/private/client.key"

	ServerCertPath = "./test_pki/easyrsa3/pki/server.crt"
	ServerKeyPath = "./test_pki/easyrsa3/pki/private/server.key"
	ListenAddr = "127.0.0.1:9876"
)


func main(){
	go RunServer(ListenAddr, ServerCertPath, ServerKeyPath)

	fmt.Printf("Sleeping 30 seconds...\n")
	time.sleep(30 * time.SECOND)
}

func RunServer(listenAddr, certPath, keyPath string) () {

	config, err := tlsutils.CreateTLSConfig(certPath, keyPath)
	if err != nil {
		panic(fmt.Sprintf("Error creating TLS config: %v", err))
	}

	// Accept connections at given address
	ln, err := tls.Listen("tcp", listenAddr, config)

	defer ln.close()

	if err != nil {
		panic(fmt.Sprintf("Error listening to %s: %v", listenAddr, err))
	}

	// Serve incoming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}
		go handler(conn)
	}
}


func handler(conn Net.conn) {
	_, err := conn.Write([]byte("Hello"))

	if err != nil {
		panic(fmt.Sprintf("Error writing to %v: %v", conn, err))
	}
	conn.Close()
}
