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
	// create a socket to listen for incoming connections
	listener, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// close the listener when the application closes
	defer listener.Close()
	fmt.Println("Listening on " + connHost + ":" + connPort)
	for {
		// accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		// handle connection in a new goroutine
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// make a buffer to hold incoming data
	buf := make([]byte, maxMsgLen)
	// read the incoming request into the buffer
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Printf("Message received. Bytes: %d \n=============\n%s", reqLen, buf)
	// send HTTP-like response back to sender
	resMsg := "Your message was delivered!"
	httpResp := fmt.Sprintf("HTTP/1.1 200 OK\nContent-Type: text/plain\nContent-Length: %d\n\n%s", len(resMsg), resMsg)
	conn.Write([]byte(httpResp))
	conn.Close()
}
