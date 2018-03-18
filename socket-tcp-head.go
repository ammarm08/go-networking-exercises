/*
  Get Head Info

  type TCPAddr struct {
    IP  IP
    Port in
  }

  <net> resolve tcp host:port address: func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)

  <net.TCPConn> duplex comms between client and server

  Write: func (c *TCPConn) Write(b []byte) (n int, err os.Error)
  Read: func (c *TCPConn) Read(b []byte) (n int, err os.Error)

  Dial: func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn)
  -> establishes a tcp conn between two addresses

  Example case:
  "HEAD / HTTP/1.0\r\n\r\n"

*/

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// establish connect b/t self and target address
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// write a HEAD HTTP request
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// consume the response
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
