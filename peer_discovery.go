package main

import "time"

type Peer struct {
	Addr     string
	LastSeen int64
}

type PeerDiscovery struct {
	Peers map[string]*Peer
}

func NewPeerDiscovery() *PeerDiscovery {
	return &PeerDiscovery{Peers: make(map[string]*Peer)}
}

func (pd *PeerDiscovery) AddPeer(addr string) {
	pd.Peers[addr] = &Peer{Addr: addr, LastSeen: time.Now().Unix()}
}

func (pd *PeerDiscovery) AlivePeers() []string {
	var list []string
	now := time.Now().Unix()
	for addr, p := range pd.Peers {
		if now-p.LastSeen < 30 {
			list = append(list, addr)
		}
	}
	return list
}
