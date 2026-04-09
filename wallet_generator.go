package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Address    string
}

func NewWallet() *Wallet {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pub := append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...)
	addr := GetPublicAddress(pub)
	return &Wallet{
		PrivateKey: priv,
		PublicKey:  pub,
		Address:    addr,
	}
}

func GetPublicAddress(pub []byte) string {
	hash := sha256.Sum256(pub)
	return hex.EncodeToString(hash[:])[:40]
}
