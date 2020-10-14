package main

import (
	"fmt"
	"simpleredis.task/clientapi"
	"os"
)

func main() {

	// access command line parameters
	cmd := os.Args[1]
	params := os.Args[2:]

	// call the api to contact the server
	var response = clientapi.ExecuteCmd(cmd, params)
	fmt.Println(response)

}
