package main

import (
	"fmt"
	"math/rand"
	"os"
)

type CloseError struct {
	ErrorDescription string
}

func (e CloseError) Error() string {
	return e.ErrorDescription
}

type Resource struct {
	Id int
}

func (r Resource) Close() error {
	_, err := fmt.Printf("Closing %v\n", r)
	if err != nil {
		return CloseError{"Close error: " + err.Error()}
	}
	if rand.Intn(2) > 0 {
		return CloseError{"Random close error"}
	}
	return nil
}

func main() {
	// Create guard
	var guard *Resource
	defer func() {
		if guard != nil {
			// No way to handle closing error :(
			_ = guard.Close()
		}
	}()
	// Create resource and immediately protect it with guard
	resource := Resource{42}
	guard = &resource
	// Some usage of resource goes here
	fmt.Printf("Created %v\n", resource)
	//panic("An error happened")
	// Normal closing of resource with immediate un-protection
	err := resource.Close()
	guard = nil
	// Handling error during resource closing
	if err != nil {
		fmt.Printf("Error during closing %v: %s\n", resource, err.Error())
		os.Exit(1)
	}
}
