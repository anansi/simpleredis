package main

import (
	"api"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 4096)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Read error: ", err)
		}

		responseData := api.HandleIncomingNetworkRequest(buffer)
		_, err = conn.Write(responseData)

		if err != nil {
			fmt.Println("Write error: ", err)
			return
		}
	}

}

func main() {

	fmt.Println("Started the server")

	// TODO implement a more friendly datastore for other developers to find understandable

	ln, err := net.Listen("tcp", ":5566")
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
