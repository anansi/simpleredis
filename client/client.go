package main

import (
	"fmt"
	"simpleredis.task/clientapi"
	"os"
	"log"
)

func main() {

	// Allow command line usage, check for valid arguments
	if len(os.Args) == 1	{
		log.Println("Error: Provided at least 1 cmd line argument: SET or GET")
		os.Exit(1)
	}

	// access command line parameter strings
	var cmd string = os.Args[1]
	var params []string = os.Args[2:]
	
	// call the clientapi to contact the server
	var response = clientapi.ExecuteCmd(cmd, params)
	fmt.Println(response)

}
