package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.In(1)
	q.In(2)
	if q.Empty() {
		t.Error("empty error")
		return
	}
	v := q.Out().(int)
	if v != 1 {
		t.Error("pop error")
	}
}
