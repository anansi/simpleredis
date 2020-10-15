package main

import (
	"fmt"
	"net"
	"simpleredis.task/serverapi"
)

const (
	networkProtocol = "tcp"
	port = ":5566"
)


func main() {

	fmt.Println("Started the server")

	// Accept concurrent network requests from clients
	ln, err := net.Listen(networkProtocol, port)
	if err != nil {
		// handle error
		fmt.Println("// handle error for ln")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println("// handle error for conn")
		}
		go handleConnection(conn)
	}

}

// handleConnection takes a network request and processes it
func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 4096)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			return
		}

		responseData := serverapi.HandleIncomingNetworkRequest(buffer)
		_, err = conn.Write(responseData)

		if err != nil {
			return
		}
	}

}