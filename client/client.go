package main

import (
	"bufio"
	"fmt"
	"local/protocol"
	"net"
)

func main() {

	fmt.Println("Started the client")

	message := "hi there"
	// call the protocol to encode message into bin
	var bin = protocol.Encode(message)
	fmt.Println(bin)
	return

	conn, err := net.Dial("tcp", "localhost:5566")
	if err != nil {
		// handle error
	}

	fmt.Fprintf(conn, message)
	status, err := bufio.NewReader(conn).ReadString('\n')
	// ...
	fmt.Println(status)

}
