package main

import (
	"fmt"
	"os"
	"strconv"
)

var m = map[int]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

func main() {
	fmt.Print("Enter integer number within [0, 10]: ")
	var text string
	_, err := fmt.Scanln(&text)
	if err != nil {
		fmt.Printf("Failed to get input: %v\n", err)
		os.Exit(1)
	}
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Printf("Failed to convert input to integer number: %v\n", err)
		os.Exit(1)
	}
	if num < 0 || num >= len(m) {
		fmt.Printf("Provided number is out of range: %v\n", num)
		os.Exit(1)
	}
	fmt.Printf("You entered: %v\n", m[num])
}
