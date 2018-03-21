/*
  TLS ECHO CLIENT
*/

package main

import (
	"crypto/tls"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := tls.Dial("tcp", service, nil)
	checkError(err)

	for n := 0; n < 5; n++ {
		// write
		fmt.Println("Writing...")
		conn.Write([]byte("Hello " + string(n+48)))

		// read res
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		checkError(err)

		// echo
		fmt.Println(string(buf[0:n]))
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
