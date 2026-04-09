package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Height     uint64
	Hash       string
	PreHash    string
	Data       string
	Timestamp  int64
	MerkleRoot string
	Nonce      int
	Difficulty int
}

func CalculateBlockHash(block Block) string {
	record := strconv.FormatUint(block.Height, 10) +
		block.PreHash +
		block.Data +
		strconv.FormatInt(block.Timestamp, 10) +
		block.MerkleRoot +
		strconv.Itoa(block.Nonce) +
		strconv.Itoa(block.Difficulty)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func GenerateBlock(preBlock Block, data string) Block {
	newBlock := Block{
		Height:    preBlock.Height + 1,
		PreHash:   preBlock.Hash,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
	newH := CalculateMerkleRoot([]string{data})
	newBlock.MerkleRoot = newH
	newBlock.Hash = CalculateBlockHash(newBlock)
	return newBlock
}
