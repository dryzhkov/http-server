package main

import (
	"fmt"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "5555"
	connType = "tcp"
)

func main() {
	// open a new connection to server
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	// send a message and close the connection
	conn.Write([]byte("New message from client\n"))
	conn.Close()
	fmt.Println("Message sent")
}
