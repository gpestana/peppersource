package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	pps "github.com/gpestana/peppersource"
	"io/ioutil"
	"log"
)

func main() {

	conf, err := getConf()
	if err != nil {
		log.Fatal(err)
	}

	c, err := pps.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	c.Run()

	// keep running
	for {
	}
}

func getConf() (pps.ClientConf, error) {
	c := pps.ClientConf{}
	cp := flag.String("conf", "", "path for configuration file")
	flag.Parse()

	if *cp == "" {
		return c, errors.New("No configuration file provided (-conf)")
	}

	craw, err := ioutil.ReadFile(*cp)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(craw, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

func printInit(c *pps.Client) {
	fmt.Println(":fire: Pepersource :fire:\n")
	for _, s := range c.Subscriptions {
		fmt.Println(s)
	}
}
