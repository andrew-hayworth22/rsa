package rsa

import (
	"math/big"
	"math/rand/v2"
)

type PublicKey struct {
	M int
	E int
}

type PrivateKey struct {
	M int
	D int
}

type KeyPair struct {
	PrivateKey PrivateKey
	PublicKey  PublicKey
}

func NewKeyPair(firstPrime int, secondPrime int) *KeyPair {
	mod := firstPrime * secondPrime
	euler := (firstPrime - 1) * (secondPrime - 1)

	enc := rand.IntN(euler)
	for gcd(enc, euler) != 1 {
		enc = rand.IntN(euler)
	}

	pub := PublicKey{
		M: mod,
		E: enc,
	}

	modInv := modInverse(enc, euler)

	priv := PrivateKey{
		M: mod,
		D: modInv,
	}

	return &KeyPair{
		PublicKey:  pub,
		PrivateKey: priv,
	}
}

func (pk *PublicKey) Encrypt(plaintext string) []int {
	plaintextBytes := []byte(plaintext)
	ciphertext := make([]int, len(plaintext))

	for idx, plaintextByte := range plaintextBytes {
		bigPlaintext := big.NewInt(int64(plaintextByte))
		bigE := big.NewInt(int64(pk.E))
		bigM := big.NewInt(int64(pk.M))

		bigPow := bigPlaintext.Exp(bigPlaintext, bigE, nil)
		bigMod := bigPow.Mod(bigPow, bigM)
		val := bigMod.Int64()

		ciphertext[idx] = int(val)
	}

	return ciphertext
}

func (pk *PrivateKey) Decrypt(ciphertext []int) string {
	plaintextBytes := make([]byte, len(ciphertext))

	for idx, ciphertextInt := range ciphertext {
		bigCiphertext := big.NewInt(int64(ciphertextInt))
		bigD := big.NewInt(int64(pk.D))
		bigM := big.NewInt(int64(pk.M))

		bigPow := bigCiphertext.Exp(bigCiphertext, bigD, nil)
		bigMod := bigPow.Mod(bigPow, bigM)
		val := bigMod.Int64()

		plaintextBytes[idx] = byte(val)
	}

	return string(plaintextBytes)
}
