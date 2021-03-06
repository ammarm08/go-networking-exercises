/*
  LookupHost

  <net> host lookup: func LookupHost(name string) (addrs []string, err os.Error)
  <net> CNAME lookup: func LookupCNAME(name string) (cname string, err os.Error)

*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}
