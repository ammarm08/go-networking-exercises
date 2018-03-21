/*
  Serialize JSON

  <json> NewEncoder(w io.Writer)
  <Encoder> Encode(v interface{})
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

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	saveJSON("person.json", person)
}

func saveJSON(fname string, key interface{}) {
	// touch
	out, err := os.Create(fname)
	checkError(err)

	// write
	encoder := json.NewEncoder(out)
	err = encoder.Encode(key)
	checkError(err)

	// close
	out.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
