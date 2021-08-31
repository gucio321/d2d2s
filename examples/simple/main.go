// Package main:
// example - simple
// description:
//	create a new character save file
package main

import (
	"io/ioutil"
	"log"

	"github.com/gucio321/d2d2s/pkg/d2s"
)

func main() {
	// create a new character
	char := d2s.New()
	char.Name = "example"

	// encode
	data, err := char.Encode()
	if err != nil {
		log.Fatal(err)
	}

	// save to file
	ioutil.WriteFile("./example.d2s", data, 0o600)
}
