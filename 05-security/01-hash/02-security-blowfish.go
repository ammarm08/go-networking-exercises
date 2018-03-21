/*
  Blowfish

  symmetric key encryption

  block algorithm (meaning if data not aligned w/ block size, then must pad blanks)

  blowfish needs 8-byte blocks

*/

package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/blowfish"
)

func main() {
	key := []byte("some key")

	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 8 byte block
	src := []byte("hello\n\n\n")
	var enc [512]byte

	// write encrypted bytes to "enc"
	cipher.Encrypt(enc[0:], src)

	// write decrypted bytes to "decrypt"
	var decrypt [8]byte
	cipher.Decrypt(decrypt[0:], enc[0:])

	// write to byte array, then coerce to string
	result := bytes.NewBuffer(nil)
	result.Write(decrypt[0:8])
	fmt.Println(string(result.Bytes()))
}
