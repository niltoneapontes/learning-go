package main

import (
	"fmt"
	"greetings/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Gladys")
	// If there IS an error, print it to console and exit program
	if err != nil {
		log.Fatal(err)
	}

	//	Else, just print the message to the console
	fmt.Println(message)
}
