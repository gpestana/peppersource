package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	conf, err := getConf()
	if err != nil {
		log.Fatal(err)
	}

	c, err := NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(c)

	c.Run()
}

type Configuration struct {
	Channels             []string
	Destination_dir      string
	Provider_pubkey_path string
}

func getConf() (Configuration, error) {
	c := Configuration{}
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

func printInit(c *Client) {
	fmt.Println(":fire: Pepersource :fire:\n")
	for _, s := range c.subscriptions {
		fmt.Println(s)
	}
}
