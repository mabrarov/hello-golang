package main

import (
	"fmt"
	"sync"
)

type ProducerId string

func produce(id ProducerId, n int, c chan int) {
	fmt.Printf("Started producer: %v\n", id)
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)
	fmt.Printf("Completed producer: %v\n", id)
}

func consume(id1, id2 ProducerId, c1, c2 chan int) {
	f1 := func(v int) {
		fmt.Printf("Received from producer %v: %d\n", id1, v)
	}
	f2 := func(v int) {
		fmt.Printf("Received from producer %v: %d\n", id2, v)
	}
	fmt.Println("Started consumer")
	for count := 0; count < 2; {
		select {
		case v, ok := <-c1:
			if ok {
				f1(v)
			} else {
				c1, f1 = c2, f2
				count++
			}
		case v, ok := <-c2:
			if ok {
				f2(v)
			} else {
				c2, f2 = c1, f1
				count++
			}
		}
	}
	fmt.Println("Completed consumer")
}

func main() {
	const producer1 ProducerId = "1"
	const producer2 ProducerId = "2"
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	var sg sync.WaitGroup
	sg.Add(2)
	go func() {
		sg.Done()
		sg.Wait()
		produce(producer1, 100, c1)
	}()
	go func() {
		sg.Done()
		sg.Wait()
		produce(producer2, 100, c2)
	}()
	consume(producer1, producer2, c1, c2)
}
