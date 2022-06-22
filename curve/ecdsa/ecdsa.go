package ecdsa

import (
	"github.com/tronch0/ecdsa/bigint"
	"github.com/tronch0/ecdsa/curve"
	"github.com/tronch0/ecdsa/curve/ecdsa/point"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func Sign(pk *PrivateKey, msgToSign string, defs curve.EcdsaCurve) *Signature {
	e := pk.Key
	z := bigint.HashStringToBigInt(msgToSign)
	k := bigint.GetRandomN(defs.GetN())
	r := defs.GetG().ScalarMul(k).GetX().GetNum()
	kInverse := new(big.Int).ModInverse(k, defs.GetN())

	interS := new(big.Int).Mul(r, e)
	interS = new(big.Int).Add(interS, z)
	interS = new(big.Int).Mul(interS, kInverse)
	s := new(big.Int).Mod(interS, defs.GetN())

	return &Signature{R: r, S: s}
}

func Verify(publicKey *point.Point, z *big.Int, sig *Signature, defs curve.EcdsaCurve) bool {
	sInverse := new(big.Int).ModInverse(sig.S, defs.GetN())
	u := new(big.Int).Mul(z, sInverse)
	u = new(big.Int).Mod(u, defs.GetN())

	v := new(big.Int).Mul(sig.R, sInverse)
	v = new(big.Int).Mod(v, defs.GetN())

	uG := defs.GetG().ScalarMul(u)
	vP := publicKey.ScalarMul(v)

	pointR := uG.Add(vP)

	return pointR.GetX().GetNum().Cmp(sig.R) == 0
}
