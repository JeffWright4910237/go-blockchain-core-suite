package main

func EncodeBlock(block Block) []byte {
	return []byte(block.Hash + block.PreHash + block.Data)
}
