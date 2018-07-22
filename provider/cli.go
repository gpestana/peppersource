package main

import (
	"log"
)

// TODO: CLI
func main() {
	p, _ := NewProvider()
	ch := "pepersource"
	fp := "/Users/gpestana/go/src/github.com/gpestana/pepersource/README.md"

	hash, err := p.Release(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = p.Notify(hash, ch)
	if err != nil {
		log.Fatal(err)
	}
}
