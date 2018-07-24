package peppersource

import (
	crypto "github.com/libp2p/go-libp2p-crypto"
)

type Head struct {
	metadata []byte
	binHash  string
	pk       crypto.PrivKey
	signed   []byte
}

func New(meta []byte, binHash string, pk crypto.PrivKey) ([]byte, err) {}
func (h *header) sign()
func (h *Header) MarshalJSON() ([]byte, error) {}
func (h *Header) UnmarshalJSON(b []byte) error {}
