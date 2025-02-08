package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ProducerId string

func produce(id ProducerId, n int, c chan int) {
	fmt.Printf("Started producer: %v\n", id)
	for i := 0; i < n; i++ {
		c <- rand.Intn(n)
		time.Sleep(time.Duration(rand.Intn(500-10)+10) * time.Millisecond)
	}
	close(c)
	fmt.Printf("Completed producer: %v\n", id)
}

func consume(id1, id2 ProducerId, c1, c2 chan int) {
	fmt.Println("Started consumer")
	for s := true; s; {
		select {
		case v, ok := <-c1:
			if ok {
				fmt.Printf("Received from producer %v: %d\n", id1, v)
			} else {
				s = false
			}
		case v, ok := <-c2:
			if ok {
				fmt.Printf("Received from producer %v: %d\n", id2, v)
			} else {
				c2 = c1
				id2 = id1
				s = false
			}
		}
	}
	for v := range c2 {
		fmt.Printf("Received from producer %v: %d\n", id2, v)
	}
	fmt.Println("Completed consumer")
}

const producer1 ProducerId = "1"
const producer2 ProducerId = "2"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go produce(producer1, 10, c1)
	go produce(producer2, 10, c2)
	consume(producer1, producer2, c1, c2)
}
