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
		log.Fatal("Error: Provided at least 1 cmd line argument: SET or GET")
	}
	
	// access command line parameter strings
	var cmd string = os.Args[1]
	var params []string = os.Args[2:]
	
	// call the clientapi to contact the server
	var response = clientapi.ExecuteCmd(cmd, params)
	fmt.Println(response)


	// Uncomment the code below to spam the server, for easy concurrency testing
	// Usage: pass valid arguments to the client in terminal, and observe an infinitely loop of SET and GET. Run the client like this on multiple terminals to create concurrency

	// //BEGIN SPAM TEST BLOCK
	// cnt := 0
	// for {
	// 	updatingSpamValue := string(cnt) + "__there"
	// 	// call the clientapi to contact the server
	// 	var response = clientapi.Set("hey", updatingSpamValue)
	// 	fmt.Println(response)
	// 	var response2 = clientapi.Get("hey")		
	// 	fmt.Println(response2)
	// 	cnt += 1
	// }
	// //END SPAM TEST BLOCK

}
