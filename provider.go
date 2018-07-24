package peppersource

import (
	"bytes"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	ipfs "github.com/ipfs/go-ipfs-api"
	crypto "github.com/libp2p/go-libp2p-crypto"
	"io/ioutil"
	"log"
	"os"
)

type ProviderConf struct {
	Path_bin string
	Pubsub   struct {
		Channels []string
	}
	PrivKeyPath string
	Metadata    map[string]string
}

type Provider struct {
	shell    *ipfs.Shell
	pk       crypto.PrivKey
	meta     []byte
	channels []string
}

func NewProvider(conf ProviderConf) (*Provider, error) {
	var prov Provider
	sh := ipfs.NewLocalShell()

	pkraw, err := ioutil.ReadFile(conf.PrivKeyPath)
	if err != nil {
		return &prov, err
	}

	if err != nil {
		return &prov, err
	}
	block, _ := pem.Decode(pkraw)
	pk, err := crypto.UnmarshalRsaPrivateKey(block.Bytes)
	if err != nil {
		return &prov, err
	}

	var meta []byte
	meta, err = json.Marshal(conf.Metadata)
	if err != nil {
		return &prov, err
	}

	if meta == nil {
		return &prov, errors.New("metadata field in configuration cannot be empty")
	}

	prov.meta = meta
	prov.shell = sh
	prov.pk = pk
	prov.channels = conf.Pubsub.Channels

	return &prov, nil
}

func (p *Provider) Release(bin string) (string, error) {
	// add file to ipfs
	bbytes, err := os.Open(bin)
	if err != nil {
		return "", err
	}
	binhash, err := p.shell.Add(bbytes)
	if err != nil {
		return "", err
	}

	// creates and signs Head
	h, err := NewHead(p.meta, binhash, p.pk)
	if err != nil {
		return "", err
	}

	hb, err := json.Marshal(&h)
	if err != nil {
		return "", err
	}

	// returns hash of signed metadata file uploaded to IPFS (used as a pointer
	// for clients to verify and download release)
	// TODO: Head hash must be saves somewhere for posterity
	return p.shell.Add(bytes.NewReader(hb))
}

func (p *Provider) Notify(h string) error {
	for _, ch := range p.channels {
		err := p.shell.PubSubPublish(ch, h)
		if err != nil {
			return err
		}
		log.Println(fmt.Sprintf("Channel '%v' notified with hash %v", ch, h))
	}
	return nil
}
