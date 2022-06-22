package bigint

import (
	"crypto/sha256"
	"math/big"
	"math/rand"
	"time"
)

func HashStringToBigInt(msg string) *big.Int {
	msgB := []byte(msg)
	msgHash := sha256.Sum256(msgB)
	msgHashSlice := msgHash[:]
	res := new(big.Int).SetBytes(msgHashSlice)
	return res
}

func HashBytesToBigInt(msg []byte) *big.Int {
	msgHash := sha256.Sum256(msg)
	msgHashSlice := msgHash[:]
	res := new(big.Int).SetBytes(msgHashSlice)
	return res
}

func GetRandom() *big.Int {
	//Max random value, a 130-bits integer, i.e 2^130 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
	time.Sleep(1 * time.Second)
	source := rand.NewSource(time.Now().UTC().Unix())

	res := new(big.Int).Rand(rand.New(source), max)
	return res
}

func GetRandomN(n *big.Int) *big.Int {
	source := rand.NewSource(time.Now().UTC().Unix())
	res := new(big.Int).Rand(rand.New(source), n)
	return res
}

func SwapEndianness(num *big.Int) *big.Int {
	b := num.Bytes()

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}

	res := new(big.Int).SetBytes(b)
	return res
}

func StringToBigInt(s string, base int) *big.Int {
	res, ok := new(big.Int).SetString(s, base)
	if ok == false {
		panic("StringToBigInt() failed")
	}

	return res
}

func Int64ToBigInt(n int64) *big.Int {
	return new(big.Int).SetInt64(n)
}

func Add(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

func Sub(x, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

func Mul(x, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}

func Exp(x, y, mod *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, mod)
}

func Div(x, y *big.Int) *big.Int {
	return new(big.Int).Div(x, y)
}
