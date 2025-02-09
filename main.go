package main

import (
	"fmt"
	"sync"
)

type CountDownLatch struct {
	cond  *sync.Cond
	value int
}

func NewCountDownLatch(value int) *CountDownLatch {
	return &CountDownLatch{cond: sync.NewCond(&sync.Mutex{}), value: value}
}

func (c *CountDownLatch) CountDownAndWait() {
	c.cond.L.Lock()
	if c.value > 0 {
		c.value--
	}
	if c.value == 0 {
		c.cond.Broadcast()
	} else {
		for c.value > 0 {
			c.cond.Wait()
		}
	}
	c.cond.L.Unlock()
}

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
	var c3 chan int
	var f3 func(v int)
	fmt.Println("Started consumer")
	for s := true; s; {
		select {
		case v, ok := <-c1:
			if ok {
				f1(v)
			} else {
				c3, f3, s = c2, f2, false
			}
		case v, ok := <-c2:
			if ok {
				f2(v)
			} else {
				c3, f3, s = c1, f1, false
			}
		}
	}
	for v := range c3 {
		f3(v)
	}
	fmt.Println("Completed consumer")
}

func main() {
	const producer1 ProducerId = "1"
	const producer2 ProducerId = "2"
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	sg := NewCountDownLatch(2)
	go func() {
		sg.CountDownAndWait()
		produce(producer1, 100, c1)
	}()
	go func() {
		sg.CountDownAndWait()
		produce(producer2, 100, c2)
	}()
	consume(producer1, producer2, c1, c2)
}
