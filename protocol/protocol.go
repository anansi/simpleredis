package protocol

import "fmt"



// Encode returns the required binary data to send across our network call
func Encode(message string) string {
    
    
	
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