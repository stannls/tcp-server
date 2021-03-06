package main

import (
	"fmt"
	"org.ydsh.tcpserver/client"
	"org.ydsh.tcpserver/server"
	"os"
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
		Clients: []*client.Client{},
	}
	s.Start()
}
