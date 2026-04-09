package main

type DPoSNode struct {
	Address string
	Votes   float64
}

func DPoSElectWitnesses(nodes []DPoSNode, count int) []DPoSNode {
	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[j].Votes > nodes[i].Votes {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
		}
	}
	if len(nodes) <= count {
		return nodes
	}
	return nodes[:count]
}
