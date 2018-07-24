package main

import (
	"encoding/json"
	"errors"
	"flag"
	pps "github.com/gpestana/peppersource"
	"io/ioutil"
	"log"
)

func main() {

	conf, err := getConf()
	if err != nil {
		log.Fatal(err)
	}

	p, _ := pps.NewProvider(conf)
	ch := "peppersource"
	fp := "/Users/gpestana/go/src/github.com/gpestana/peppersource/README.md"

	hash, err := p.Release(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = p.Notify(hash, ch)
	if err != nil {
		log.Fatal(err)
	}
}

func getConf() (pps.ProviderConf, error) {
	var c pps.ProviderConf
	cp := flag.String("conf", "", "path for configuration file")
	pp := flag.String("pk", "", "path for RSA provate key file")
	flag.Parse()

	if *cp == "" {
		return c, errors.New("No configuration file provided (-conf)")
	}

	if *pp == "" {
		return c, errors.New("No private key file provided (-pk)")
	}
	craw, err := ioutil.ReadFile(*cp)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(craw, &c)
	if err != nil {
		return c, err
	}

	c.PrivKeyPath = *pp

	return c, err
}
