/*
  GOB LOAD
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

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {
	var person Person
	loadGob("person.gob", &person)

	fmt.Println("Person:", person.String())
	os.Exit(0)
}

func loadGob(fname string, key interface{}) {
	in, err := os.Open(fname)
	checkError(err)

	decoder := gob.NewDecoder(in)
	err2 := decoder.Decode(key)
	checkError(err2)

	in.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
