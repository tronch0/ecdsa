package ecdsa

import (
	"github.com/tronch0/ecdsa/bigint"
	"github.com/tronch0/ecdsa/curve/ecdsa/secp256k1"
	"testing"
)

func TestPublicKeyGeneration(t *testing.T) {
	secretKey := bigint.GetRandom()
	privateKey := NewPrivateKey(secp256k1.GetSecp256k1(), secretKey)
	publicKey := privateKey.PublicKey

	g := secp256k1.GetSecp256k1().GetG()
	expectedPK := g.ScalarMul(privateKey.Key)

	if publicKey.Equal(expectedPK) == false {
		t.Fatal("the assertion has failed")
	}
}
