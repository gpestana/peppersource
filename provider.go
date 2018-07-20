package pepersource

import (
	ipfs "github.com/ipfs/go-ipfs-api"
	"os"
)

type Provider struct {
	shell *ipfs.Shell
}

func NewProvider() (*Provider, error) {
	sh := ipfs.NewLocalShell()

	return &Provider{
		shell: sh,
	}, nil
}

func (p *Provider) Release(pt string) (string, error) {

	// add file to ipfs
	f, err := os.Open(pt)
	if err != nil {
		return "", err
	}

	return p.shell.Add(f)
}

func (p *Provider) Notify(h string, ch string) error {
	err := p.shell.PubSubPublish(ch, h)
	if err != nil {
		return err
	}
	return nil
}
