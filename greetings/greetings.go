package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// A slice of message formats.
var formats = []string{
	"Hi, %v. Welcome!",
	"Great to see you, %v!",
	"Hail, %v! Well met!",
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names ...string) ([]string, error) {
	messages := make([]string, len(names))
	for i, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[i] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
