package main

import (
	"fmt"
	"net"
)

var m map[string]string

func handleConnection(conn net.Conn) {
	fmt.Println("handling listening to Connection", conn)
	defer conn.Close()
	buffer := make([]byte, 24)
	for {
		n, err := conn.Read(buffer)

		if n > 0 {
			// fmt.Println(message)
		}

		var resp_string = m["greeting"]
		_, err = conn.Write([]byte(resp_string))

		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func main() {

	fmt.Println("Started the server")

	// TODO implement a more friendly datastore for other developers to find understandable
	m = make(map[string]string)
	m["greeting"] = "howzit"

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
