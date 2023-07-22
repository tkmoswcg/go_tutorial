package greetings

import "fmt"

// hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Wlecome!", name)
	return message
}
