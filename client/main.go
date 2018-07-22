package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {

	dir, chs := cli()
	c, err := NewClient(dir, chs)
	if err != nil {
		log.Fatal(err)
	}

	c.Run()
}

func cli() (string, []string) {
	dir := flag.String("dir", ".", "workspace directory where bundles and metadata are stored")
	ch := flag.String("channels", "", "comma separated list of notification channels")
	flag.Parse()

	var chs []string
	for _, c := range strings.Split(*ch, ",") {
		chs = append(chs, c)
	}
	return *dir, chs
}

func printInit(c *Client) {
	fmt.Println(":fire: Pepersource :fire:\n")
	for _, s := range c.subscriptions {
		fmt.Println(s)
	}
}
