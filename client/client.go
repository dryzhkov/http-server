package main

import (
	"fmt"
	"net"
	"os"
)

const (
	connHost  = "localhost"
	connPort  = "5555"
	connType  = "tcp"
	maxMsgLen = 4096
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
	// read response from server
	buf := make([]byte, maxMsgLen)
	reqLen, err := conn.Read(buf)
	conn.Close()
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Printf("Message sent. Server response below. Bytes: %d \n=============\n%s\n", reqLen, buf)
}
