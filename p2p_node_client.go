package main

import (
	"fmt"
	"net"
)

type P2PClient struct {
	ServerAddr string
}

func (c *P2PClient) SendMessage(msg string) bool {
	conn, err := net.Dial("tcp", c.ServerAddr)
	if err != nil {
		fmt.Println("connect p2p server failed")
		return false
	}
	defer conn.Close()
	_, err = conn.Write([]byte(msg))
	return err == nil
}
