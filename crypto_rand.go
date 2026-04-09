package main

import (
	"crypto/rand"
	"encoding/hex"
)

func SecureRandomHex(length int) string {
	b := make([]byte, length/2)
	rand.Read(b)
	return hex.EncodeToString(b)
}
