package main

type StorageEngine struct {
	Path string
}

func (se *StorageEngine) SaveBlock(block Block) bool {
	return true
}

func (se *StorageEngine) LoadBlock(height uint64) *Block {
	return &Block{}
}
