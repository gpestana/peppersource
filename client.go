package pepersource

import (
	ipfs "github.com/ipfs/go-ipfs-api"
	"log"
)

type Client struct {
	shell         *ipfs.Shell
	subscriptions []*ipfs.PubSubSubscription
}

// subscribes to channels and returns client
func NewClient(ss ...string) (*Client, error) {
	sh := ipfs.NewLocalShell()
	var subs []*ipfs.PubSubSubscription

	for _, s := range ss {
		pbs, err := sh.PubSubSubscribe(s)
		if err != nil {
			return nil, err
		}
		subs = append(subs, pbs)
		log.Println("Client: subscribed to ", s)
	}

	return &Client{
		shell:         sh,
		subscriptions: subs,
	}, nil
}

// Client runs as a daemon waiting for notifications
func (c *Client) Run() {}
