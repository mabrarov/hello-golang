package main

import (
	"fmt"
	"github.com/mabrarov/hello-golang/greetings"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC | log.Lshortfile | log.Lmsgprefix)

	// Get a greeting message and print it.
	messages, err := greetings.Hellos("Gladys", "Samantha", "Darrin")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	for _, message := range messages {
		fmt.Println(message)
	}
}
