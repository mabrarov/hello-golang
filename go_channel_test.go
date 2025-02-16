package main

import (
	"testing"
)

func TestSendToClosedGoChannelPanic(t *testing.T) {
	var ch Channel[int] = NewGoChannel[int]()
	if !verifySendToClosedChannelPanic(ch) {
		t.Fatal("Expected to panic when sending to closed channel")
	}
}

func TestCloseOfClosedGoChannelPanic(t *testing.T) {
	var ch Channel[int] = NewGoChannel[int]()
	if !verifyCloseOfClosedChannelPanic(ch) {
		t.Fatal("Expected to panic when closing already closed channel")
	}
}

func TestNoReadAfterCloseGoChannel(t *testing.T) {
	var ch Channel[int] = NewGoChannel[int]()
	if !verifyReadAfterClose(ch, genInt) {
		t.Fatal("Successfully read from channel while previous read was unsuccessful due to channel is closed")
	}
}
