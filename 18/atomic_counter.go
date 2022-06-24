package main

import "sync/atomic"

// счетчик с использованием атомарных операций
type atomicCounter struct {
	counter int64
}

func newAtomicStruct() *atomicCounter {
	return &atomicCounter{
		counter: 0,
	}
}

func (a *atomicCounter) Add() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *atomicCounter) Get() int64 {
	return atomic.LoadInt64(&a.counter)
}
