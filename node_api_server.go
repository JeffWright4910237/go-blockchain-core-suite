package main

import "fmt"

type APIServer struct {
	Port int
}

func (api *APIServer) StartAPI() {
	fmt.Printf("API server running on port %d\n", api.Port)
}
