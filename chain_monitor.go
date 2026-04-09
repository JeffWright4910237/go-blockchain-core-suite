package main

import "time"

type ChainMonitor struct {
	LastBlockTime int64
}

func (cm *ChainMonitor) CheckBlockSpeed() int64 {
	return time.Now().Unix() - cm.LastBlockTime
}
