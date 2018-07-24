package peppersource

import (
	"encoding/pem"
	"fmt"
	ipfs "github.com/ipfs/go-ipfs-api"
	crypto "github.com/libp2p/go-libp2p-crypto"
	"io/ioutil"
	"log"
)

type ClientConf struct {
	Channels             []string
	Destination_dir      string
	Provider_pubkey_path string
}

type Client struct {
	shell         *ipfs.Shell
	Subscriptions []*Subscription
	dir           string
	provider_pk   crypto.PubKey
}

// subscribes to channels and returns client
func NewClient(conf ClientConf) (*Client, error) {
	cli := &Client{}
	sh := ipfs.NewLocalShell()
	dir := conf.Destination_dir

	// loads provider pk
	bpk, err := ioutil.ReadFile(conf.Provider_pubkey_path)
	if err != nil {
		return cli, err
	}
	block, _ := pem.Decode(bpk)
	pk, err := crypto.UnmarshalRsaPublicKey(block.Bytes)
	if err != nil {
		return cli, err
	}

	notificationHandler := func(t string, r ipfs.PubSubRecord) {
		h := string(r.Data())
		log.Println(fmt.Sprintf("pubsub (%v); %v: data { %v }", t, r.From(), h))

		// TODO: do verifications

		// ok: download
		err = cli.shell.Get(h, dir)
		if err != nil {
			log.Println("ERROR: ", err)
			return
		}
		log.Println(fmt.Sprintf("downloaded %v to %v", h, dir))
	}

	var subs []*Subscription
	for _, c := range conf.Channels {
		s := Subscription{notificationHandler: notificationHandler,
			client: cli,
			topic:  c,
			on:     false,
		}
		subs = append(subs, &s)
	}

	cli.shell = sh
	cli.dir = dir
	cli.provider_pk = pk
	cli.Subscriptions = subs

	return cli, nil
}

func Get(hash string) {}

// Client runs as a daemon waiting for notifications
func (c *Client) Run() {
	// subscribe to all channels
	for _, s := range c.Subscriptions {
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
