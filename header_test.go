package peppersource

import (
	"encoding/json"
	"fmt"
	crypto "github.com/libp2p/go-libp2p-crypto"
	"testing"
)

func TestSignVerify(t *testing.T) {
	privk, pubk, err := crypto.GenerateKeyPair(crypto.RSA, 1024)
	if err != nil {
		t.Fatal(err)
	}

	hash := "test_hash"
	metam := map[string]string{
		"meta1": "tes_1t",
		"meta2": "test_2",
	}

	metab, err := json.Marshal(metam)
	if err != nil {
		t.Fatal(err)
	}

	h, err := NewHead(metab, hash, privk)
	if err != nil {
		t.Fatal(err)
	}

	// unmarshal for sending in the wire

	pb, err := json.Marshal(h.Payload)
	if err != nil {
		t.Fatal(err)
	}

	if string(h.Metadata()) != string(metab) {
		t.Error(fmt.Sprintf("Payload metadata should be %v, got %v", string(metab), string(h.Metadata())))
	}

	if h.Hash() != hash {
		t.Error(fmt.Sprintf("Payload hash should be %v, got %v", hash, h.Hash()))
	}

	valid, err := verifyHead(pb, pubk)
	if valid != false {
		t.Error("Verification should have returned valid signature")
	}
}
