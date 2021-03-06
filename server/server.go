package server

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	Port       int
	MaxPlayers int
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
		go server.handleConnection(c)
	}
}

func (server *Server) handleConnection(c net.Conn) {

	log.Printf("Accepting connection from %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil && err.Error() == "EOF" {
			log.Printf("Client with adress %s disconnected\n", c.RemoteAddr().String())
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		result := "Hello Client\n"
		c.Write([]byte(string(result)))
	}
	log.Printf("Closing connection from %s\n", c.RemoteAddr().String())
	c.Close()
}
