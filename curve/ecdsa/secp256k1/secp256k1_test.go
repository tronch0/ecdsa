package secp256k1

import (
	"fmt"
	"github.com/tronch0/ecdsa/curve/ecdsa"
	"github.com/tronch0/ecdsa/curve/ecdsa/point"
	"github.com/tronch0/ecdsa/field"
	"math/big"
	"testing"
)

func TestVerication(t *testing.T) {

	x, ok := new(big.Int).SetString("887387e452b8eacc4acfde10d9aaf7f6d9a0f975aabb10d006e4da568744d06c", 16)
	if ok == false {
		panic("couldn't parse x to field defs")
	}

	y, ok := new(big.Int).SetString("61de6d95231cd89026e286df3b6ae4a894a3378e393e93a0f45b666329a0ae34", 16)
	if ok == false {
		panic("couldn't parse y to field defs")
	}

	p1 := point.New(
		field.New(x, GetSecp256k1().GetP()), // x
		field.New(y, GetSecp256k1().GetP()), // y
		GetSecp256k1().GetA(),               // a
		GetSecp256k1().GetB(),               // b
	)

	z, ok := new(big.Int).SetString("ec208baa0fc1c19f708a9ca96fdeff3ac3f230bb4a7ba4aede4942ad003c0f60", 16)
	if ok == false {
		panic("couldn't parse z to field defs")
	}

	r, ok := new(big.Int).SetString("ac8d1c87e51d0d441be8b3dd5b05c8795b48875dffe00b7ffcfac23010d3a395", 16)
	if ok == false {
		panic("couldn't parse r to field defs")
	}

	s, ok := new(big.Int).SetString("68342ceff8935ededd102dd876ffd6ba72d6a427a3edb13d26eb0781cb423c4", 16)
	if ok == false {
		panic("couldn't parse s to field defs")
	}

	sig := &ecdsa.Signature{R: r, S: s}
	res := ecdsa.Verify(p1, z, sig, GetSecp256k1())
	fmt.Println(res)

}

// to update parameters
func TestVerication2(t *testing.T) {

	x, ok := new(big.Int).SetString("887387e452b8eacc4acfde10d9aaf7f6d9a0f975aabb10d006e4da568744d06c", 16)
	if ok == false {
		panic("couldn't parse x to field defs")
	}

	y, ok := new(big.Int).SetString("61de6d95231cd89026e286df3b6ae4a894a3378e393e93a0f45b666329a0ae34", 16)
	if ok == false {
		panic("couldn't parse y to field defs")
	}

	p1 := point.New(
		field.New(x, GetSecp256k1().GetP()), // x
		field.New(y, GetSecp256k1().GetP()), // y
		GetSecp256k1().GetA(),               // a
		GetSecp256k1().GetB(),               // b
	)

	z, ok := new(big.Int).SetString("ec208baa0fc1c19f708a9ca96fdeff3ac3f230bb4a7ba4aede4942ad003c0f60", 16)
	if ok == false {
		panic("couldn't parse z to field defs")
	}

	r, ok := new(big.Int).SetString("ac8d1c87e51d0d441be8b3dd5b05c8795b48875dffe00b7ffcfac23010d3a395", 16)
	if ok == false {
		panic("couldn't parse r to field defs")
	}

	s, ok := new(big.Int).SetString("68342ceff8935ededd102dd876ffd6ba72d6a427a3edb13d26eb0781cb423c4", 16)
	if ok == false {
		panic("couldn't parse s to field defs")
	}

	sig := &ecdsa.Signature{R: r, S: s}
	res := ecdsa.Verify(p1, z, sig, GetSecp256k1())
	fmt.Println(res)

}
