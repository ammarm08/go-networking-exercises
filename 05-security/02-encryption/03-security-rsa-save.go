/*
  RSA

  public key encryption (requiring two keys)
  - one to encrypt (public) asymmmetric
  - one to decrypt (private)
*/

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	reader := rand.Reader
	bitSize := 512

	// generate a 512-bit rsa key
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := key.PublicKey
	fmt.Println("Public key modulus", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)

	saveGobKey("private.key", key)
	saveGobKey("public.key", publicKey)

	savePEMKey("private.pem", key)
}

func saveGobKey(fname string, key interface{}) {
	out, err := os.Create(fname)
	checkError(err)

	encoder := gob.NewEncoder(out)
	err = encoder.Encode(key)
	checkError(err)

	out.Close()
}

func savePEMKey(fname string, key *rsa.PrivateKey) {
	out, err := os.Create(fname)
	checkError(err)

	var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key)}

	pem.Encode(out, privateKey)

	out.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Fatal error %", err.Error())
		os.Exit(1)
	}
}
