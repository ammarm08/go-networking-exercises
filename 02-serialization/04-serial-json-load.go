/*
  Deserialize JSON

  <json> NewDecoder(r io.Reader)
  <Decoder> Decode (k interface{})
*/

package main

import (
	"encoding/json"
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

// method on Person
func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {
	var person Person

	// pass reference to decoder target
	loadJSON("person.json", &person)

	fmt.Println("Person", person.String())
}

func loadJSON(fname string, key interface{}) {
	// open
	in, err := os.Open(fname)
	checkError(err)

	// decode
	decoder := json.NewDecoder(in)
	err = decoder.Decode(key)
	checkError(err)

  // close
	in.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
