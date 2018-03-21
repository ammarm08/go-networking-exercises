/*
  UTF8 CLIENT
*/

package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)

// byte order marking (big endian)
const BOM = '\ufffe'

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	// read from connection (returns []uint16)
	shorts := readShorts(conn)

	// decode to list of int32
	ints := utf16.Decode(shorts)

	// coerce to string
	str := string(ints)

	fmt.Println("[]int16", shorts, len(shorts))
	fmt.Println("decoded to []int32", ints, len(ints))
	fmt.Println("decoded to string:", str, len(str))

	os.Exit(0)
}

func readShorts(conn net.Conn) []uint16 {
	var buf [512]byte

	// first two bytes be endianness
	n, err := conn.Read(buf[0:2])

	for true {
		// read into buffer until done
		m, err := conn.Read(buf[n:])
		if m == 0 || err != nil {
			break
		}

		n += m
	}

	checkError(err)

	// need only half the space for int16 vs int32
	var shorts []uint16
	shorts = make([]uint16, n/2)

	if buf[0] == 0xff && buf[1] == 0xfe {
		// big endian
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i])<<8 + uint16(buf[i+1])
		}
	} else if buf[0] == 0xfe && buf[1] == 0xff {
		// little endian
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i]) + uint16(buf[i+1])<<8
		}
	} else {
		fmt.Println("Unknown order")
	}

	return shorts
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
