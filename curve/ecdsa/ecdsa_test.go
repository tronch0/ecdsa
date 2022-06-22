package ecdsa

import (
	"github.com/tronch0/ecdsa/bigint"
	"github.com/tronch0/ecdsa/curve/ecdsa/P256"
	"github.com/tronch0/ecdsa/curve/ecdsa/secp256k1"
	"testing"
)

func TestFullFlow_Secp256k1(t *testing.T) {
	secretKey := bigint.GetRandom()
	privateKey := NewPrivateKey(secp256k1.GetSecp256k1(), secretKey)

	// sign
	msgToSign := "Heyy"
	sig := Sign(privateKey, msgToSign, secp256k1.GetSecp256k1())

	// verify
	z := bigint.HashStringToBigInt(msgToSign)
	verificationRes := Verify(privateKey.PublicKey, z, sig, secp256k1.GetSecp256k1())

	if verificationRes == false {
		t.Fatal("the verification of the signature has failed")
	}
}

func TestFullFlow_P256(t *testing.T) {
	secretKey := bigint.GetRandom()
	privateKey := NewPrivateKey(P256.GetP256(), secretKey)

	// sign
	msgToSign := "Heyy"
	sig := Sign(privateKey, msgToSign, P256.GetP256())

	// verify
	z := bigint.HashStringToBigInt(msgToSign)
	verificationRes := Verify(privateKey.PublicKey, z, sig, P256.GetP256())

	if verificationRes == false {
		t.Fatal("the verification of the signature has failed")
	}
}

func TestDiffPublicKeyVerification(t *testing.T) {
	secretKey := bigint.GetRandom()
	privateKey := NewPrivateKey(secp256k1.GetSecp256k1(), secretKey)

	// sign
	msgToSign := "Heyy"
	sig := Sign(privateKey, msgToSign, secp256k1.GetSecp256k1())

	// verify
	z := bigint.HashStringToBigInt(msgToSign)
	secondSecretKey := bigint.GetRandom()
	wrongPrivateKey := NewPrivateKey(secp256k1.GetSecp256k1(), secondSecretKey)
	verificationRes := Verify(wrongPrivateKey.PublicKey, z, sig, secp256k1.GetSecp256k1())
	if verificationRes {
		t.Fatal("succeed to verify signature (verification expected to fail)")
	}
}
