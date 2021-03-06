package server

import (
	"github.com/google/uuid"
	"log"
	"net"
	"sort"
	"strconv"
	"tcpserver/client"
)

type Server struct {
	Port       int
	MaxPlayers int
	Clients    []*client.Client
}

func (server *Server) Start() {
	l, err := net.Listen("tcp4", ":"+strconv.Itoa(server.Port))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Serving Gameserver on Port %d\n", server.Port)
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		if len(server.Clients) == server.MaxPlayers{
			c.Write([]byte("Server full\n"))
			c.Close()
		} else {
			go server.handleConnection(c)
		}
	}
}

func (server *Server) handleConnection(c net.Conn) {
	log.Printf("Accepting connection from %s\n", c.RemoteAddr().String())
	client := client.Client{Connection: c, Id: uuid.New().String()}
	server.Clients = append(server.Clients, &client)
	var welcome = "Hello Client\n"
	client.SendMessage(&welcome)
	for {
		message, err := client.ReceiveMessage()
		if err != nil {
			log.Printf("Client with adress %s disconnected\n", client.Connection.RemoteAddr().String())
			break
		}
		log.Printf("Client wrote %s", message)
	}
	client.Disconnect()
	server.RemoveClient(&client)
}

func (server *Server) RemoveClient(client *client.Client) {
	i := sort.Search(len(server.Clients), func(i int) bool {
		return server.Clients[i] == client
	})
	if len(server.Clients) > 1 {
		server.Clients = append(server.Clients[:i], server.Clients[i+1:]...)
	} else {
		server.Clients = nil
	}
}
