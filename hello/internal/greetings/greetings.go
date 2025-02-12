package greetings

import (
	"github.com/mabrarov/hello-golang/greetings"
)

func GenGreetings() ([]string, error) {
	return greetings.Hellos("Gladys", "Samantha", "Darrin")
}
