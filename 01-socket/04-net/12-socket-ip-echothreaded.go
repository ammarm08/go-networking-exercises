/*
  ThreadedIPEchoServer

  <net> Listen : func Listen(net, laddr string) (l Listener, err os.Error)
  <*Addr> Accept : func Accept() (c Conn, err os.Error)
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":1200"
	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

// this time, the conn is of type net.Conn (generic)
func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
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
