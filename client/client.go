package main

import (
	"bufio"
	"fmt"
	"local/api"
	"net"
)

func main() {

	fmt.Println("Started the client")

	keyToGet := "greeting"
	// call the protocol to encode message into bin
	var bin = api.Get(keyToGet)
	fmt.Println(bin)
	return

	conn, err := net.Dial("tcp", "localhost:5566")
	if err != nil {
		// handle error
	}

	fmt.Fprintf(conn, "TODO implement networking via protocol")
	status, err := bufio.NewReader(conn).ReadString('\n')
	// ...
	fmt.Println(status)

}
