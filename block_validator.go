package main

func ValidateBlock(newBlock, preBlock Block) bool {
	if newBlock.Height != preBlock.Height+1 {
		return false
	}
	if newBlock.PreHash != preBlock.Hash {
		return false
	}
	if CalculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
