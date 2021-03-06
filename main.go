package main

import (
	"log"
	"os"
	"strconv"
	"tcpserver/client"
	"tcpserver/server"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatal("Please provide a port number!")
		return
	}
	port, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Panic("Invalid Port number")
	}
	s := server.Server{
		Port:       port,
		MaxPlayers: 100,
		Clients:    []*client.Client{},
	}
	s.Start()
}
