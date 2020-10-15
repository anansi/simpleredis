package clientapi

import (
	"log"
	"net"
	"strings"
	"simpleredis.task/coder"
)

const (
	get = "GET"
	set = "SET"
	serverIpPort = "localhost:5566"
	networkProtocol = "tcp"
)

// ExeuteCmd allows a client to execute any command supplied with its parameters.
// It returns the relivant response from the API call 
func ExecuteCmd(cmd string, params []string) string {

	switch strings.ToUpper(cmd) { // ToUpper makes the user input case insensitive
	case "GET":
		if len(params) != 1	{
			log.Fatalf("Error: GET cmd requires 1 parameter, but %d were provided. exiting", len(params))
		}
		return Get(params[0])
	case "SET":
		if len(params) != 2	{
			log.Fatal("Error: SET cmd requires 2 parameters, but ",len(params)," parameter(s) were provided. exiting")
		}
		return Set(params[0], params[1])
	
	}

	log.Fatal("Error: Provided unsupported command: ", strings.ToUpper(cmd))
	return ""// Go requires me to end this function with a return. 
}

// Get allows a Client to ask the server to retreive the value of a key in the simpleredis datastore
// Get could be accessed by the client via the ExecuteCmd function above
// Get is made Public, for convinience for the client.
func Get(key string) string {

	// TODO encode the words for the relivant Get command
	getWord := coder.Encode(get)
	keyWord := coder.Encode(key)

	concat := append(getWord, keyWord...)

	// TODO send the command to the server
	// TODO get the response from the server
	responseData := sendNetworkRequest(concat)
	var response []string = coder.Decode(responseData)

	// TODO return the value obtained, as required
	if len(response) == 0	{
		// a known response from the README.md protocol, is an empty string (no word)
		return ""
	}	else	{
		return response[0]
	}
}

// Set lets the client set a new value in the simpleredis datastore
// Set is made Public, for convinience for the client
func Set(key, value string) string {

	// TODO encode the words for the relivant Set command
	setWord := coder.Encode(set)
	keyWord := coder.Encode(key)
	valueWord := coder.Encode(value)

	concat := append(append(setWord, keyWord...), valueWord...)

	// TODO send the command to the server
	// TODO get the response from the server
	response := sendNetworkRequest(concat)

	// TODO return the value obtained, as required
	strResponse := coder.Decode(response)[0]
	return strResponse

}

// sendNetworkRequest is a private function for sending data to the server
func sendNetworkRequest(data []byte) []byte {
	conn, err := net.Dial(networkProtocol, serverIpPort)
	if err != nil {
		// handle error
		log.Fatalf("Error: Server is not running, is it running at %s?",serverIpPort)
	}
	defer conn.Close() // final call of this function

	// write the data to the connection
	_, err = conn.Write(data)

	replyData := make([]byte, 2048)
	conn.Read(replyData)

	return replyData

}
