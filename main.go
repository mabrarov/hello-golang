package main

import (
	"fmt"
)

type Resource struct {
	Id int
}

func (r Resource) Close() {
	fmt.Printf("Closing %v\n", r)
}

func main() {
	var r *Resource
	defer func() {
		if r != nil {
			r.Close()
		}
	}()
	r = &Resource{42}
	fmt.Printf("Created %v\n", *r)
	//panic("An error happened")
	r.Close()
	r = nil
}
