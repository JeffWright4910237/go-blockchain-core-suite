package main

import "time"

func GetBlockTimestamp() int64 {
	return time.Now().Unix()
}
