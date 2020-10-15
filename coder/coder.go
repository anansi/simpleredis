package coder

import	(
	"encoding/binary"
)

// Encode returns the required binary data for a word to send across our network call
// the meaning of 'word' is: a space separated keyword used in our API
// Note that byte is an alias for uint8
func Encode(word string) []byte {

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


func Decode(data []byte) []string	{

	// Note words include GET, SET, and relivant parameters
	var words []string = parseDataForWords(data)

	return words
}

func parseDataForWords(data []byte) []string {

	var cmdParameterList []string = make([]string, 0)

	nextByteIndex := uint16(0)
	// loop until there are no more words to parse. This is when the expected word length parameter is 00
	for {

		// get the next word from the data. We know each word data starts with a 2 byte length field
		nextWordLengthData := data[nextByteIndex : nextByteIndex+2]
		wordLength := binary.BigEndian.Uint16(nextWordLengthData)

		// if wordLength is 00, the data has stopped
		if wordLength == 0 {
			// break out of the for loop, no more parameters according to protocol
			break
		}

		nextWordStartIndex := nextByteIndex + 2
		nextWordEndIndex := nextWordStartIndex + wordLength
		nextWordData := data[nextWordStartIndex:nextWordEndIndex]
		nextWordValue := string(nextWordData)

		// add the new word to the slice to return
		cmdParameterList = append(cmdParameterList, nextWordValue)

		nextByteIndex = nextWordStartIndex + wordLength // the end of the range for the next word

	}

	return cmdParameterList
}

