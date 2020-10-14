package clientapi

import (
	"net"
	"strings"
	"simpleredis.task/coder"
)

const get = "GET"
const set = "SET"

// TODO implement this properly
func sendNetworkRequest(data []byte) string {
	conn, err := net.Dial("tcp", "localhost:5566")
	defer conn.Close()

	if err != nil {
		// handle error
	}

	// write the data to the connection
	_, err = conn.Write(data)

	reply := make([]byte, 2048)
	conn.Read(reply)

	return string(reply)

}

// ExeuteCmd allows a client to execute any command supplied with its parameters.
// It returns the relivant response from the API call or protocol command
func ExecuteCmd(cmd string, params []string) string {

	switch strings.ToUpper(cmd) { // ToUpper makes the user input
	case "GET":
		return Get(params[0])
	case "SET":
		return Set(params[0], params[1])
	}

	return "provided unsupported command"
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
	response := sendNetworkRequest(concat)

	// TODO return the value obtained, as required
	return response

}

// Set lets the client set a value in the simpleredis datastore
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
	return response

}
