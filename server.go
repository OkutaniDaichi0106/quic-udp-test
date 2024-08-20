package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/quic-go/quic-go"
)

func main() {
	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: 8443})
	if err != nil {
		fmt.Print(err.Error())
	}
	tr := quic.Transport{
		Conn: udpConn,
	}
	tlsCert, err := tls.LoadX509KeyPair("localhost.pem", "localhost-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	tlsConf := tls.Config{
		Certificates: []tls.Certificate{tlsCert},
	}
	ln, err := tr.Listen(&tlsConf, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	for {
		conn, err := ln.Accept(context.Background())
		if err != nil {
			fmt.Print(err.Error())
		}
		if conn == nil {
			continue
		}
		go Stream(conn)
	}
}

func Stream(conn quic.Connection) {
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		fmt.Print(err.Error())
	}

	io.Copy(stream, stream)
}
