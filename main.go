package main

import "fmt"

func main() {
	ch := NewCustomChannel[int]()
	go func() {
		for i := 0; i < 10; i++ {
			ch.Send(i)
		}
		_ = ch.Close()
	}()
	for v, ok := ch.Receive(); ok; v, ok = ch.Receive() {
		fmt.Println(v)
	}
}
