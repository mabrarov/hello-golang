package main

import (
	"sync"
)

type CustomChannel[T any] struct {
	cond   *sync.Cond
	value  T
	ready  bool
	closed bool
}

func NewCustomChannel[T any]() *CustomChannel[T] {
	return &CustomChannel[T]{cond: sync.NewCond(&sync.Mutex{})}
}

func (ch *CustomChannel[T]) Send(value T) {
	ch.cond.L.Lock()
	defer ch.cond.L.Unlock()

	if ch.closed {
		panic("channel already closed")
	}

	for ch.ready {
		ch.cond.Wait()
	}

	ch.value = value
	ch.ready = true
	ch.cond.Broadcast()

	for ch.ready {
		ch.cond.Wait()
	}
}

func (ch *CustomChannel[T]) Receive() (value T, ok bool) {
	ch.cond.L.Lock()
	defer ch.cond.L.Unlock()

	for !ch.ready && !ch.closed {
		ch.cond.Wait()
	}

	if !ch.ready {
		var zero T
		value, ok = zero, false
		return
	}

	value, ok = ch.value, true
	ch.ready = false
	ch.cond.Broadcast()
	return
}

func (ch *CustomChannel[T]) Close() error {
	ch.cond.L.Lock()
	defer ch.cond.L.Unlock()

	if ch.closed {
		panic("channel already closed")
	}

	ch.closed = true
	ch.cond.Broadcast()
	return nil
}
