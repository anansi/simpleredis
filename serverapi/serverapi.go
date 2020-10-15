package serverapi

import (
	"simpleredis.task/protocol"
	"simpleredis.task/coder"
)

const get = "GET"
const set = "SET"

// The server side function to handle an incoming request
func HandleIncomingNetworkRequest(data []byte) []byte {

	var words []string = coder.Decode(data)
	// get protocol input variables
	cmd := words[0]
	params := words[1:]
	var datastoreResponse string = protocol.ExecuteCmd(cmd, params)
	// encode the response string for the network,
	var datastoreData = coder.Encode(datastoreResponse)
	// send the data back to the client
	return datastoreData
}
