/*
  Encoding/ASN.1
  Abstract Syntax Notation One (1984 telecom standard)

  Marshal : func Marshal(val interface{}) ([]byte, err os.Error)
  Unmarshal : func Unmarshal(val interface{}, b []byte) (rest []byte, err os.Error)
*/

package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

func main() {
	// serialize the integer 13
	mdata, err := asn1.Marshal(13)
	checkError(err)

	// declare int n
	var n int

	// deserialize to the n-addressed int
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)

	fmt.Println("Post marshal/unmarshal: ", n)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatail error %s", err.Error())
		os.Exit(1)
	}
}
