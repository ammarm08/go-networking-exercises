/*
  Client: Daytime (over UDP)
  ** in conjunction w/ socket-udp-daytime-server **

  <net> Resolve : func ResolveUDPAddr(net, addr string) (*UDPAddr, os.Error)
  <net> DialUDP
  <net> ListenUDP
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	// write to server
	_, err = conn.Write([]byte("anything"))
	checkError(err)

	// consume reply
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
