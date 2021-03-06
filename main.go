package main

import (
	"fmt"
	"os"
	"tcpserver/client"
	"tcpserver/server"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	s := server.Server{
		Port:       8000,
		MaxPlayers: 100,
		Clients:    []*client.Client{},
	}
	s.Start()
}
