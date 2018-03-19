/*
  Mask: type IPMask []byte

  create mask: func IPv4Mask(a, b, c, d byte) IPMask
  default mask: func (ip IP) DefaultMask() IPMask
  find network from mask: func (ip IP) Mask(mask IPMask) IP

*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	mask := addr.DefaultMask() // bc addr is type IP
	network := addr.Mask(mask) // method on type IP that takes an IPMask
	ones, bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		"Default mask length is ", bits,
		"Leading ones count is ", ones,
		"Mask is (hex) ", mask.String(),
		"Network is ", network.String())

	os.Exit(0)
}
