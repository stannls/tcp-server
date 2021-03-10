package server

import (
	"github.com/google/uuid"
	"log"
	"net"
	"strconv"
	"tcpserver/client"
	"tcpserver/events"
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
		if len(server.Clients) == server.MaxPlayers {
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
	server.Broadcast(events.SpawnEvent{PlayerId: client.Id}.ToJson(), client.Id)
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
	server.Broadcast(events.DisconnectEvent{PlayerId: client.Id}.ToJson(), client.Id)
}

func (server *Server) RemoveClient(client *client.Client) {
	var position int = 0
	for i, c := range server.Clients {
		if c.Id == client.Id {
			position = i
			break
		}
	}
	if len(server.Clients) > 1 {
		server.Clients = append(server.Clients[:position], server.Clients[position+1:]...)
	} else {
		server.Clients = nil
	}
}

func (server *Server) Broadcast(data string, sender string) {
	for _, c := range server.Clients {
		if c.Id != sender {
			_, _ = c.Connection.Write([]byte(data))
		}
	}
}
