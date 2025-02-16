package main

type Channel[T any] interface {
	Send(value T)
	Receive() (T, bool)
	Close() error
}
