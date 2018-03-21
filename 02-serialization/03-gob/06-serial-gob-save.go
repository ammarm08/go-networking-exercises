/*
GOB SAVE

gob: serialization specific to Go

*/

package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	saveGob("person.gob", person)
}

func saveGob(fname string, key interface{}) {
	out, err := os.Create(fname)
	checkError(err)

	encoder := gob.NewEncoder(out)
	err2 := encoder.Encode(key)
	checkError(err2)

	out.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
