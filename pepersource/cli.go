package main

import (
	ppsrc "github.com/gpestana/pepersource"
	"log"
)

// TODO: CLI
func main() {
	c, err := ppsrc.NewClient("pepersource", "any", "thing")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(c)

	p, _ := ppsrc.NewProvider()
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
