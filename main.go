package main

import "fmt"

func main() {
	m := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
	}
	p := m[1]
	fmt.Println(p)
}
