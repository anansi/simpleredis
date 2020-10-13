package protocol

import "fmt"



// Encode returns the required binary data to send across our network call
func Encode(message string) string {
    
    // determine the length of the message, as this is the first 2 bytes
    length := len(message)
    fmt.Println("length: ", length)
    length_bytes := fmt.Sprintf("%08b", byte(length))
    fmt.Println("length_bytes: ", length_bytes)
    // initiaise the output binary value as the length of the message, in binary
    // output := ""
	
	fmt.Println("about to Encode: ", message)

	for i := 0; i < len(message); i++ {
		fmt.Println(i, ": ", message[i], "(decimal)")

		char := message[i]
		byte := fmt.Sprintf("%08b", byte(char))
		fmt.Println(i, ": ", byte, "(byte)")
	}
	


    result := fmt.Sprintf("Hi, %v. Welcome!", message)
    return result
}