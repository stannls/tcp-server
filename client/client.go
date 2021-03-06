package client

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type Client struct {
	Connection net.Conn
	Id         string
}

func (client *Client) ReceiveMessage() (string, error) {
	message, err := bufio.NewReader(client.Connection).ReadString('\n')
	if err != nil && err.Error() == "EOF" {
		return "", err
	}
	return strings.TrimSpace(string(message)), nil
}

func (client *Client) Disconnect() {
	log.Printf("Closing connection from %s\n", client.Connection.RemoteAddr().String())
	_ = client.Connection.Close()
}

func (client *Client) SendMessage(message *string) {
	_, err := client.Connection.Write([]byte(*message))
	if err != nil {
		log.Fatal(err)
	}
}
