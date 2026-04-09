package main

import (
	"fmt"
	"net"
	"time"
)

type P2PServer struct {
	Addr    string
	NodeID  string
	IsAlive bool
}

func NewP2PServer(addr string) *P2PServer {
	return &P2PServer{
		Addr:    addr,
		NodeID:  SHA256Hash(addr + time.Now().String()),
		IsAlive: true,
	}
}

func (s *P2PServer) Start() {
	listen, err := net.Listen("tcp", s.Addr)
	if err != nil {
		fmt.Println("P2P server start failed")
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			break
		}
		go s.HandleConn(conn)
	}
}

func (s *P2PServer) HandleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	conn.Read(buf)
	fmt.Println("receive p2p message:", string(buf))
}
