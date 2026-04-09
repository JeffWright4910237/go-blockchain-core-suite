package main

import (
	"strconv"
	"strings"
)

func PoWMining(block Block, difficulty int) (int, string) {
	target := strings.Repeat("0", difficulty)
	nonce := 0
	for {
		block.Nonce = nonce
		hash := CalculateBlockHash(block)
		if strings.HasPrefix(hash, target) {
			return nonce, hash
		}
		nonce++
	}
}
