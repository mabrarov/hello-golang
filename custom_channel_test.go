package main

import (
	"sync"
	"testing"
)

func TestSendToClosedCustomChannelPanic(t *testing.T) {
	var ch Channel[int] = NewCustomChannel[int]()
	if !verifySendToClosedChannelPanic(ch) {
		t.Fatal("Expected to panic when sending to closed channel")
	}
}

func TestCloseOfClosedCustomChannelPanic(t *testing.T) {
	var ch Channel[int] = NewCustomChannel[int]()
	if !verifyCloseOfClosedChannelPanic(ch) {
		t.Fatal("Expected to panic when closing already closed channel")
	}
}

func TestReceiveFromClosedCustomChannel(t *testing.T) {
	ch := NewCustomChannel[int]()
	_ = ch.Close()
	v, ok := ch.Receive()
	if v != 0 || ok {
		t.Fatalf("Expected to receive from closed channel: 0, false. Got: %d, %v", v, ok)
	}
}

func TestSingleSenderSingleReceiverCustomChannel(t *testing.T) {
	const n = 10
	ch := NewCustomChannel[int]()
	go func() {
		for i := 0; i < n; i++ {
			ch.Send(i)
		}
		_ = ch.Close()
	}()
	for i := 0; i < n; i++ {
		v, ok := ch.Receive()
		if !ok {
			t.Fatal("Expected successful read from channel, but got closed channel")
		}
		if v != i {
			t.Fatalf("Expected to read: %d, got: %d", i, v)
		}
	}
	_, ok := ch.Receive()
	if ok {
		t.Fatal("Expected false when reading from closed channel, but got true")
	}
}

func TestMultipleSendersSingleReceiverCustomChannel(t *testing.T) {
	const n = 10
	const senders = 10
	ch := NewCustomChannel[int]()

	c := sync.OnceFunc(func() {
		_ = ch.Close()
	})
	var sent sync.WaitGroup
	sent.Add(senders)
	for i := 0; i < senders; i++ {
		go func() {
			defer c()
			for i := 0; i < n; i++ {
				ch.Send(i)
			}
			sent.Done()
			sent.Wait()
		}()
	}

	m := make(map[int]int)
	for i := 0; i < n*senders; i++ {
		v, ok := ch.Receive()
		if !ok {
			t.Fatal("Expected successful read from channel, but got closed channel")
		}
		m[v]++
	}

	_, ok := ch.Receive()
	if ok {
		t.Fatal("Expected false when reading from closed channel, but got true")
	}
	for i, v := range m {
		if v != n {
			t.Errorf("Expected to receive %d value %d time(s), but got %d time(s).", i, n, v)
		}
	}
}

func TestSingleSenderMultipleReceiversCustomChannel(t *testing.T) {
	const n = 10
	const receivers = 10
	ch := NewCustomChannel[int]()

	m := make(map[int]map[int]int)
	var received sync.WaitGroup
	received.Add(receivers)
	for i := 0; i < receivers; i++ {
		r := make(map[int]int)
		m[i] = r
		go func() {
			defer received.Done()
			for v, ok := ch.Receive(); ok; v, ok = ch.Receive() {
				r[v]++
			}
		}()
	}

	for i := 0; i < n; i++ {
		for j := 0; j < receivers; j++ {
			ch.Send(i)
		}
	}
	_ = ch.Close()

	received.Wait()

	total := make(map[int]int)
	for _, v := range m {
		for i, c := range v {
			total[i] += c
		}
	}

	for i, v := range total {
		if v != n {
			t.Errorf("Expected to receive %d value %d time(s), but got %d time(s).", i, n, v)
		}
	}
}

func TestNoReadAfterCloseCustomChannel(t *testing.T) {
	var ch Channel[int] = NewCustomChannel[int]()
	if !verifyReadAfterClose(ch, genInt) {
		t.Fatal("Successfully read from channel while previous read was unsuccessful")
	}
}
