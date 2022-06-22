package secp256k1

import (
	"github.com/tronch0/ecdsa/curve"
	"github.com/tronch0/ecdsa/curve/ecdsa/point"
	"github.com/tronch0/ecdsa/field"
	"math/big"
)

var secp256k1Defs curve.EcdsaCurve

func GetSecp256k1() curve.EcdsaCurve {
	if secp256k1Defs == nil {
		secp256k1Defs = newSecp256k1()
	}

	return secp256k1Defs
}

func newSecp256k1() *Secp256k1 {
	N, ok := new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	if ok == false {
		panic("couldn't parse n to field defs")
	}

	p, ok := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
	if ok == false {
		panic("couldn't parse prime to field defs")
	}

	gX, ok := new(big.Int).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
	if ok == false {
		panic("couldn't parse g.gX to field defs")
	}

	gY, ok := new(big.Int).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)
	if ok == false {
		panic("couldn't parse g.y to field defs")
	}

	a := big.NewInt(0)
	b := big.NewInt(7)

	ffA := field.New(a, p)
	ffB := field.New(b, p)
	x := field.New(gX, p)
	y := field.New(gY, p)

	gP := point.New(x, y, ffA, ffB)

	return &Secp256k1{
		n:      N,
		p:      p,
		a:      field.New(a, p),
		b:      field.New(b, p),
		gPoint: gP,
	}
}

type Secp256k1 struct {
	n      *big.Int
	p      *big.Int
	a      *field.Element
	b      *field.Element
	gPoint *point.Point
}

func (s *Secp256k1) GetG() *point.Point {
	return s.gPoint
}

func (s *Secp256k1) GetA() *field.Element {
	return s.a
}

func (s *Secp256k1) GetB() *field.Element {
	return s.b

}

func (s *Secp256k1) GetP() *big.Int {
	return s.p
}

func (s *Secp256k1) GetN() *big.Int {
	return s.n
}
