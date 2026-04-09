package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type BlockChain struct {
	Chain           []Block
	CurrentHeight   uint64
	GenesisHash     string
	ConsensusType   string
}

func InitBlockChain() *BlockChain {
	genesis := CreateGenesisBlock()
	chain := []Block{genesis}
	return &BlockChain{
		Chain:         chain,
		CurrentHeight: 1,
		GenesisHash:   genesis.Hash,
		ConsensusType: "pow",
	}
}

func (bc *BlockChain) AddBlockToChain(block Block) bool {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	if block.PreHash == lastBlock.Hash && block.Height == lastBlock.Height+1 {
		bc.Chain = append(bc.Chain, block)
		bc.CurrentHeight++
		return true
	}
	return false
}

func (bc *BlockChain) GetChainHash() string {
	data := ""
	for _, b := range bc.Chain {
		data += b.Hash
	}
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
