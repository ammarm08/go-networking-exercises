/*
  TLS Echo Server

  TLS (transport layer security, formerly SSL)
  * supports encrypted message passing
  * takes care a lot of the encryption/decryption for you

  client-server negotiate identity using x.509 certificates
  upon which a secret key is created just for their use
*/

package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("jan.newmarch.name.pem", "private.pem")
	checkError(err)

	config := tls.Config{Certificates: []tls.Certificate{cert}}

	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader

	service := "0.0.0.0:1200"

	listener, err := tls.Listen("tcp", service, &config)
	checkError(err)

	fmt.Println("Listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("Accepted")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		fmt.Println("Trying to read")
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
