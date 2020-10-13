package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Println("handling Connection")
}

func main() {

	fmt.Println("Hello, World!")

	ln, err := net.Listen("tcp", ":8080")
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
