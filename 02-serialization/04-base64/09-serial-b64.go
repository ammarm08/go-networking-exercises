/*
Base64

base64.NewEncoder(enc *Encoding, w io.Writer) io WriteCloser
base64.NewDecoder(enc *Encoding, r io.Reader) io.Reader
*/

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	eightBytes := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	// a Buffer reference (which is an io Writer/Reader)
	bb := &bytes.Buffer{}

	// b64 encode to bb
	encoder := base64.NewEncoder(base64.StdEncoding, bb)
	encoder.Write(eightBytes)
	encoder.Close()

	fmt.Println(bb)

	// 12-byte buffer
	dbuf := make([]byte, 12)
	decoder := base64.NewDecoder(base64.StdEncoding, bb)
	decoder.Read(dbuf)

	for _, ch := range dbuf {
		fmt.Print(ch)
	}
}
