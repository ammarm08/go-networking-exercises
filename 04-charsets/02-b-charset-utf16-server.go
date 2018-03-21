/*
  UTF16 Server
*/

package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)

// Byte Order Marker (FFFE = BE, FEFF = LE)
const BOM = '\ufffe'

func main() {
	service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		str := "arrêté"

		// encode to list of int32s (aka "runes")
		shorts := utf16.Encode([]rune(str))

		// write int stream to conn
		writeShorts(conn, shorts)

		conn.Close()
	}
}

func writeShorts(conn net.Conn, shorts []uint16) {
	var bytes [2]byte

	// establish byte order
	bytes[0] = BOM >> 8
	bytes[1] = BOM & 255
	_, err := conn.Write(bytes[0:])
	if err != nil {
		return
	}

	for _, v := range shorts {
		// write each rune as 2 fixed bytes
		bytes[0] = byte(v >> 8)
		bytes[1] = byte(v & 255)

		_, err = conn.Write(bytes[0:])
		if err != nil {
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
