package P256

import (
	"github.com/tronch0/ecdsa/curve"
	"github.com/tronch0/ecdsa/curve/ecdsa/point"
	"github.com/tronch0/ecdsa/field"
	"math/big"
)

var P256Defs curve.EcdsaCurve

func GetP256() curve.EcdsaCurve {
	if P256Defs == nil {
		P256Defs = newP256Defs()
	}

	return P256Defs
}

func newP256Defs() *P256 {
	N, ok := new(big.Int).SetString("ffffffff00000000ffffffffffffffffbce6faada7179e84f3b9cac2fc632551", 16)
	if ok == false {
		panic("couldn't parse curve parameter n")
	}

	p, ok := new(big.Int).SetString("ffffffff00000001000000000000000000000000ffffffffffffffffffffffff", 16)
	if ok == false {
		panic("couldn't parse curve parameter p")
	}

	gX, ok := new(big.Int).SetString("6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c296", 16)
	if ok == false {
		panic("couldn't parse curve parameter g.x")
	}

	gY, ok := new(big.Int).SetString("4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5", 16)
	if ok == false {
		panic("couldn't parse curve parameter g.y")
	}

	a, ok := new(big.Int).SetString("ffffffff00000001000000000000000000000000fffffffffffffffffffffffc", 16)
	if ok == false {
		panic("couldn't parse curve parameter a")
	}

	b, ok := new(big.Int).SetString("5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b", 16)
	if ok == false {
		panic("couldn't parse curve parameter b")
	}

	ffA := field.New(a, p)
	ffB := field.New(b, p)
	x := field.New(gX, p)
	y := field.New(gY, p)

	gP := point.New(x, y, ffA, ffB)

	return &P256{
		n:      N,
		p:      p,
		a:      field.New(a, p),
		b:      field.New(b, p),
		gPoint: gP,
	}
}

type P256 struct {
	n      *big.Int
	p      *big.Int
	a      *field.Element
	b      *field.Element
	gPoint *point.Point
}

func (s *P256) GetG() *point.Point {
	return s.gPoint
}

func (s *P256) GetA() *field.Element {
	return s.a
}

func (s *P256) GetB() *field.Element {
	return s.b
}

func (s *P256) GetP() *big.Int {
	return s.p
}

func (s *P256) GetN() *big.Int {
	return s.n
}
