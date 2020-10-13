package protocol

import "fmt"

// Encode returns the required binary data to send across our network call
func Encode(message string) string {
    // Return a greeting that embeds the name in a message.
    result := fmt.Sprintf("Hi, %v. Welcome!", message)
    return result
}