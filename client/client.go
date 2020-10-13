package main

import (
	"fmt"
	"net"
	"bufio"
	"local/protocol"
)


func main() {

	fmt.Println("Started the client")

	conn, err := net.Dial("tcp", "localhost:5566")
	if err != nil {
		// handle error
	}
 

	message := "hello"
	// call the protocol to encode message into bin
	protocol.Encode(message)

	fmt.Fprintf(conn, message)
	status, err := bufio.NewReader(conn).ReadString('\n')
	// ...
	fmt.Println(status)

}
