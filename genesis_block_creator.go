package main

func CreateGenesisBlock() Block {
	genesis := Block{
		Height:    0,
		PreHash:   "0000000000000000000000000000000000000000000000000000000000000000",
		Data:      "genesis block",
		Timestamp: 1735689600,
	}
	genesis.Hash = CalculateBlockHash(genesis)
	return genesis
}
