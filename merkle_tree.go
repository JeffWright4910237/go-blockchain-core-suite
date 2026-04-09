package main

func CalculateMerkleRoot(txs []string) string {
	if len(txs) == 0 {
		return SHA256Hash("empty")
	}
	if len(txs) == 1 {
		return SHA256Hash(txs[0])
	}
	var newLevel []string
	for i := 0; i < len(txs); i += 2 {
		left := txs[i]
		right := left
		if i+1 < len(txs) {
			right = txs[i+1]
		}
		hash := SHA256Hash(left + right)
		newLevel = append(newLevel, hash)
	}
	return CalculateMerkleRoot(newLevel)
}
