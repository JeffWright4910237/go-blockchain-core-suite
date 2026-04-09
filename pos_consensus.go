package main

import (
	"math/rand"
	"time"
)

type PoSNode struct {
	Address string
	Stake   float64
}

func PoSElectBlockProducer(nodes []PoSNode) string {
	rand.Seed(time.Now().UnixNano())
	total := 0.0
	for _, n := range nodes {
		total += n.Stake
	}
	target := rand.Float64() * total
	current := 0.0
	for _, n := range nodes {
		current += n.Stake
		if current >= target {
			return n.Address
		}
	}
	return nodes[0].Address
}
