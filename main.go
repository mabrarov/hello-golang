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
	fmt.Println("Started consumer")
	for c1 != nil || c2 != nil {
		select {
		case v, ok := <-c1:
			if ok {
				fmt.Printf("Received from producer %v: %d\n", id1, v)
			} else {
				c1 = nil
			}
		case v, ok := <-c2:
			if ok {
				fmt.Printf("Received from producer %v: %d\n", id2, v)
			} else {
				c2 = nil
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
