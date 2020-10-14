package main

import (
	"fmt"
	"local/api"
	"os"
)

func main() {

	// access command line parameters
	cmd := os.Args[1]
	params := os.Args[2:]

	// call the api to contact the server
	var response = api.ExecuteCmd(cmd, params)
	fmt.Println(response)

}
