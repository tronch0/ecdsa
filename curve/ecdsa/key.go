package ecdsa

import (
	"github.com/tronch0/ecdsa/curve"
	"github.com/tronch0/ecdsa/curve/ecdsa/point"
	"math/big"
)

type PrivateKey struct {
	Curve     curve.EcdsaCurve `json:"curve"`
	Key       *big.Int         `json:"key"`
	PublicKey *point.Point     `json:"public_key"`
}

func NewPrivateKey(defs curve.EcdsaCurve, k *big.Int) *PrivateKey {
	gP := defs.GetG()
	k = new(big.Int).Mod(k, defs.GetN())

	return &PrivateKey{
		Curve:     defs,
		PublicKey: gP.ScalarMul(k),
		Key:       k,
	}
}
