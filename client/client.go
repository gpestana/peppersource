package main

import (
	"fmt"
	ipfs "github.com/ipfs/go-ipfs-api"
	"log"
)

type Client struct {
	shell         *ipfs.Shell
	subscriptions []*Subscription
	dir           string
}

// subscribes to channels and returns client
func NewClient(out string, ch []string) (*Client, error) {
	sh := ipfs.NewLocalShell()

	cli := &Client{
		shell: sh,
		dir:   out,
	}

	notificationHandler := func(t string, r ipfs.PubSubRecord) {
		h := string(r.Data())
		log.Println(fmt.Sprintf("pubsub (%v); %v: data { %v }", t, r.From(), h))
		// TODO: do verifications
		// ok: download
		err := cli.shell.Get(h, out)
		if err != nil {
			log.Println("ERROR: ", err)
			return
		}
		log.Println(fmt.Sprintf("downloaded %v to %v", h, out))
	}

	var subs []*Subscription
	for _, c := range ch {
		s := Subscription{notificationHandler: notificationHandler,
			client: cli,
			topic:  c,
			on:     false,
		}
		subs = append(subs, &s)
	}

	cli.subscriptions = subs
	return cli, nil
}

func Get(hash string) {}

// Client runs as a daemon waiting for notifications
func (c *Client) Run() {
	// subscribe to all channels
	for _, s := range c.subscriptions {
		err := s.Subscribe()
		if err == nil {
			s.on = true
		} else {
			log.Println(fmt.Sprintf("failed to subscribe to %s: %s", s.topic, err))
		}
	}
	// run as deamon
	for {
	}
}

type Subscription struct {
	notificationHandler func(string, ipfs.PubSubRecord)
	client              *Client
	topic               string
	pubsub              *ipfs.PubSubSubscription
	err                 error
	on                  bool
}

func (s *Subscription) Subscribe() error {
	pbs, err := s.client.shell.PubSubSubscribe(s.topic)
	if err != nil {
		s.err = err
		return err
	}
	s.pubsub = pbs

	go s.listen()
	return nil
}

// starts listen to subscripton. once receives data from the channel, starts
// verification and download process and restart listening the same channel
func (s *Subscription) listen() {
	rec, err := s.pubsub.Next()
	if err != nil {
		log.Println(err)
		return
	}

	s.notificationHandler(s.topic, rec)
	go s.listen()
}

func (s *Subscription) String() string {
	if s.err != nil {
		return fmt.Sprintf("subscription: %s; ON: %v; err:%v", s.topic, s.on, s.err)
	}
	return fmt.Sprintf("subscription: %s; ON: %v", s.topic, s.on)
}
