/*
  Ping (ICMP protocol) using raw sockets
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host", os.Args[0])
		os.Exit(1)
	}

	addr, err := net.ResolveIPAddr("ip", os.Args[1])
	checkError(err)

	conn, err := net.DialIP("ip4:icmp", addr, addr)
	checkError(err)

	// icmp protocol
	var msg [512]byte
	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum
	msg[3] = 0  // checksum
	msg[4] = 0  // id[0]
	msg[5] = 13 // id[1]
	msg[6] = 0  // seq[0]
	msg[7] = 37 // seq[1]
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	// write from buf
	_, err = conn.Write(msg[0:len])
	checkError(err)

	// read into buf
	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("Got response")
	if msg[5] == 13 {
		fmt.Println("identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("sequence matches")
	}

	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0

	// sum based on alternating message bytes
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}

	// some shifting/masking
	sum = (sum >> 16) + (sum & 0xfff)
	sum += (sum >> 16)

	// return NOT sum
	var answer uint16 = uint16(^sum)
	return answer
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
