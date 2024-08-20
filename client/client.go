package main

import (
	"context"
	"crypto/tls"
	"log"

	"github.com/quic-go/quic-go"
)

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	tlsConf := tls.Config{InsecureSkipVerify: true}
	conn, err := quic.DialAddr(context.Background(), "localhost:8443", &tlsConf, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.CloseWithError(0, "Client closing")

	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer stream.Close()

	// Send a message
	message := "hello"

	stream.Write([]byte(message))

	// Receive a message
	buf := make([]byte, len(message))
	stream.Read(buf)

	log.Print(string(buf))
}
