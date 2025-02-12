package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellosWithoutNames(t *testing.T) {
	msg, err := Hellos()
	if len(msg) > 0 || err != nil {
		t.Fatalf(`Hellos() = %q, %v, want [], nil`, msg, err)
	}
}

func TestHellosEmptyName(t *testing.T) {
	msg, err := Hellos("")
	if len(msg) > 0 || err == nil {
		t.Fatalf(`Hellos("") = %q, %v, want [], error`, msg, err)
	}
}

func TestHellosMultipleNames(t *testing.T) {
	names := []string{"One", "Two", "Free"}
	msg, err := Hellos(names...)
	if len(msg) != 3 || err != nil {
		t.Fatalf(`Hellos("") = %q, %v, want [3]string, nil`, msg, err)
	}
	for i, name := range names {
		want := regexp.MustCompile(`\b` + name + `\b`)
		if !want.MatchString(msg[i]) {
			t.Fatalf(`Got %v, want match for %#q`, msg, want)
		}
	}
}
