package main

import (
	"sync"
	"sync/atomic"
)

func sendToChannel[T any](ch Channel[T], v T) (p any) {
	defer func() {
		p = recover()
	}()
	ch.Send(v)
	return
}

func closeChannel[T any](ch Channel[T]) (p any) {
	defer func() {
		p = recover()
	}()
	_ = ch.Close()
	return
}

func genInt(i int) int {
	return i
}

func verifySendToClosedChannelPanic[T any](ch Channel[T]) bool {
	var z T
	closeChannel(ch)
	return sendToChannel(ch, z) != nil
}

func verifyCloseOfClosedChannelPanic[T any](ch Channel[T]) bool {
	closeChannel(ch)
	return closeChannel(ch) != nil
}

func verifyReadAfterClose[T any](ch Channel[T], gen func(int) T) bool {
	const n = 1000
	const senders = 1000
	const receivers = 1000

	var closed atomic.Bool
	var readAfterClose atomic.Bool

	var received sync.WaitGroup
	received.Add(receivers)
	for i := 0; i < receivers; i++ {
		go func() {
			defer received.Done()
			for {
				if readAfterClose.Load() {
					break
				}
				done := closed.Load()
				_, ok := ch.Receive()
				if !ok {
					closed.Store(true)
					break
				}
				if done && ok {
					readAfterClose.Store(true)
					break
				}
			}
		}()
	}

	for i := 0; i < senders; i++ {
		go func() {
			for i := 0; i < n; i++ {
				sendToChannel(ch, gen(i))
			}
			closeChannel(ch)
		}()
	}

	received.Wait()

	return !readAfterClose.Load()
}
