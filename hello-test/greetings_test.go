package hello_test

import (
	"github.com/mabrarov/hello-golang/hello/internal/greetings"
	_ "github.com/testcontainers/testcontainers-go/modules/compose"
	"testing"
)

func TestExample(t *testing.T) {
	_, _ = greetings.GenGreetings()
}
