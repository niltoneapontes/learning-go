package greetings

import "fmt"

// Hello() receives a name (string) and returns a greeting message
func Hello(name string) string {
	// Returns a greeting!
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
