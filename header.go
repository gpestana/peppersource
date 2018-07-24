package peppersource

import (
	"encoding/json"
	crypto "github.com/libp2p/go-libp2p-crypto"
)

type Head struct {
	Payload   Payload
	Signature []byte
}

type Payload struct {
	Metadata string
	Hash     string
}

// NewHead returns a new Head object which contains sign(metadata, hash)
func NewHead(meta []byte, hash string, pk crypto.PrivKey) (Head, error) {
	var h Head
	p := Payload{
		Metadata: string(meta),
		Hash:     hash,
	}

	pb, err := json.Marshal(p)
	if err != nil {
		return h, err
	}

	s, err := pk.Sign(pb)
	if err != nil {
		return h, err
	}

	h.Payload = p
	h.Signature = s

	return h, nil
}

func (h Head) Hash() string {
	return h.Payload.Hash
}

func (h Head) Metadata() []byte {
	return []byte(h.Payload.Metadata)
}

// verifyHead receives a byte encoded Head and verifies if Head was signed by
// the expected entity
func verifyHead(b []byte, pubk crypto.PubKey) (bool, error) {
	var h Head
	err := json.Unmarshal(b, &h)
	if err != nil {
		return false, err
	}

	pb, err := json.Marshal(h.Payload)
	if err != nil {
		return false, err
	}

	return pubk.Verify(pb, h.Signature)
}
