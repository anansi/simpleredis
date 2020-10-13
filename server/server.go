package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Println("handling Connection", conn)
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		message := string(buffer[:n])

		if message == "/quit" {
			fmt.Println("quit command received. Bye.")
			return
		}

		if n > 0 {
			fmt.Println(message)
		}

		if err != nil {
			log.Println(err)
			return
		}
	}

}

func main() {

	fmt.Println("Started the server")

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
