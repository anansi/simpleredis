package main

import (
	"fmt"
	"local/api"
)

func main() {

	fmt.Println("Started the client")

	keyToGet := "greeting"
	// call the protocol to encode message into bin
	var bin = api.Get(keyToGet)
	fmt.Println(bin)

}
