package main

type GoChannel[T any] struct {
	ch chan T
}

func NewGoChannel[T any]() *GoChannel[T] {
	return &GoChannel[T]{make(chan T)}
}

func (ch *GoChannel[T]) Send(value T) {
	ch.ch <- value
}

func (ch *GoChannel[T]) Receive() (value T, ok bool) {
	value, ok = <-ch.ch
	return
}

func (ch *GoChannel[T]) Close() error {
	close(ch.ch)
	return nil
}
