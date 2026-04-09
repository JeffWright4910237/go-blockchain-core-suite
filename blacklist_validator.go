package main

type BlackList struct {
	Addrs map[string]bool
}

func NewBlackList() *BlackList {
	return &BlackList{Addrs: make(map[string]bool)}
}

func (bl *BlackList) IsForbidden(addr string) bool {
	return bl.Addrs[addr]
}
