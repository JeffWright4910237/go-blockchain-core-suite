package main

func ChooseBestChain(chains [][]Block) []Block {
	best := chains[0]
	for _, c := range chains {
		if len(c) > len(best) {
			best = c
		}
	}
	return best
}
