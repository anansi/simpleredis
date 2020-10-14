package serverapi

import (
	"encoding/binary"
	"simpleredis.task/protocol"
)

const get = "GET"
const set = "SET"

// Encode returns the required binary data for a word to send across our network call
// the meaning of 'word' is: a space separated keyword used in our API
// Note that byte is an alias for uint8
func encode(word string) []byte {

	// determine the length of the word, as this is the first 2 bytes
	wordLength := len(word)
	packetLength := wordLength + 2

	var outputBytes = make([]byte, packetLength)
	outputBytes[0], outputBytes[1] = uint8(wordLength>>8), uint8(wordLength&0xff)

	for i := 0; i < len(word); i++ {

		char := word[i]
		outputBytes[i+2] = byte(char) // the index is +2 as the first 2 bytes are for the word length

	}

	return outputBytes
}


func parseCommandFromData(data []byte) (string, uint16) {

	// get the length of the initial command token (first 2 bytes)
	protocolCmdLengthBytes := data[:2]
	protocolCmdLength := binary.BigEndian.Uint16(protocolCmdLengthBytes)

	// get the command token from the data
	commandBytes := data[2 : 2+protocolCmdLength]

	return string(commandBytes), protocolCmdLength

}

func parseParameterData(data []byte) []string {

	var cmdParameterList []string = make([]string, 0)

	nextByteIndex := uint16(0)
	// loop until there are no more token to parse. This is when the expected token length parameter is 00
	for {

		// get the next token from the data. We know each token starts with a 2 byte length field
		nextTokenLengthData := data[nextByteIndex : nextByteIndex+2]
		tokenLength := binary.BigEndian.Uint16(nextTokenLengthData)

		// if tokenLength is 00, the data has stopped
		if tokenLength == 0 {
			// break out of the for loop, no more parameters according to protocol
			break
		}

		nextTokenStartIndex := nextByteIndex + 2
		nextTokenEndIndex := nextTokenStartIndex + tokenLength
		nextTokenData := data[nextTokenStartIndex:nextTokenEndIndex]
		nextTokenValue := string(nextTokenData)

		// add the new token to the slice to return
		cmdParameterList = append(cmdParameterList, nextTokenValue)

		nextByteIndex += nextTokenStartIndex + tokenLength // the end of the range for the next token

	}

	return cmdParameterList
}

// The server side function to handle an incoming request
func HandleIncomingNetworkRequest(data []byte) []byte {

	// the data variables and their types
	var protocolCmd string
	var cmdLength uint16
	// determine the command being requested from the front of the data buffer
	protocolCmd, cmdLength = parseCommandFromData(data)

	// create a new slice and parse the rest of the buffer as the expected parameters to the protocol command
	var parameterData = data[cmdLength+2:]
	var parameters []string = parseParameterData(parameterData)

	// get protocol response (as a string)
	var datastoreResponse string = protocol.ExecuteCmd(protocolCmd, parameters)
	// encode the response string for the network,
	var datastoreData = encode(datastoreResponse)
	// send the data back to the client
	return datastoreData
}
