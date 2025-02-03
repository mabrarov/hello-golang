package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	dict := make(map[string]int)
	last := len(s) - 1
	var word string
	for i, c := range s {
		boundary := i == last
		switch c {
		case ' ', '\t', '\r', '\n':
			boundary = true
		default:
			word += string(c)
		}
		if boundary && len(word) > 0 {
			dict[word] += 1
			word = ""
		}
	}
	return dict
}

func main() {
	wc.Test(WordCount)
}
