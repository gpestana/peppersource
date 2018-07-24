package peppersource

import (
	"bytes"
	"encoding/json"
	"encoding/pem"
	"errors"
	ipfs "github.com/ipfs/go-ipfs-api"
	crypto "github.com/libp2p/go-libp2p-crypto"
	"io"
	"io/ioutil"
	"os"
)

type ProviderConf struct {
	Path_bin string
	Pubsub   struct {
		channels []string
	}
	PrivKeyPath string
	Metadata    interface{}
}

type Provider struct {
	shell *ipfs.Shell
	pk    crypto.PrivKey
	meta  []byte
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

	// create signed release_meta file and add it to IPFS
	mbytes, err := buildMeta(binhash, p.meta, p.pk)
	if err != nil {
		return "", err
	}

	// returns hash of signed metadata file uploaded to IPFS (used as a pointer
	// for clients to verify and download release)
	return p.shell.Add(mbytes)
}

func (p *Provider) Notify(h string, ch string) error {
	err := p.shell.PubSubPublish(ch, h)
	if err != nil {
		return err
	}
	return nil
}

func buildMeta(bhash string, meta []byte, pk crypto.PrivKey) (io.Reader, error) {
	var mr io.Reader
	mr = bytes.NewReader(meta)
	return mr, nil
}
