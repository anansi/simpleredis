package api

import (
	"bufio"
	"fmt"
	"net"
)

// note byte is an alias for uint8

// Encode returns the required binary data for a word to send across our network call
// the meaning of 'word' is: a space separated keyword used in our API
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

// TODO implement this properly
func sendNetworkRequest() {

	conn, err := net.Dial("tcp", "localhost:5566")
	if err != nil {
		// handle error
	}

	fmt.Fprintf(conn, "TODO implement networking via protocol")
	status, err := bufio.NewReader(conn).ReadString('\n')
	// ...
	fmt.Println(status)

}

// Get allows a Client to ask the server to retreive the value of a key in the simpleredis datastore
func Get(key string) string {

	// TODO encode the words for the relivant Get command

	// TODO send the command to the server

	// TODO get the response from the server

	// TODO return the value obtained, as required
	return "Get value for " + key

}
