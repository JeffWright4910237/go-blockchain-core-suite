package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

func ECDSASign(priv *ecdsa.PrivateKey, data string) (r, s *big.Int, err error) {
	hash := sha256.Sum256([]byte(data))
	r, s, err = ecdsa.Sign(rand.Reader, priv, hash[:])
	return
}

func ECDSAVerify(pub *ecdsa.PublicKey, data string, r, s *big.Int) bool {
	hash := sha256.Sum256([]byte(data))
	return ecdsa.Verify(pub, hash[:], r, s)
}
