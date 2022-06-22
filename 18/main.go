package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter interface {
	Add()
	Get() int64
}

type mutexStr struct {
	counter int64
	mutex   *sync.RWMutex
}

func newMutexStruct() *mutexStr {
	return &mutexStr{
		counter: 0,
		mutex:   &sync.RWMutex{},
	}
}

func (m *mutexStr) Add() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.counter++
}

func (m *mutexStr) Get() int64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.counter
}

type atomicStruct struct {
	counter int64
}

func newAtomicStruct() *atomicStruct {
	return &atomicStruct{
		counter: 0,
	}
}

func (a *atomicStruct) Add() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *atomicStruct) Get() int64 {
	return atomic.LoadInt64(&a.counter)
}

var workersNum = 20
var iterationsNum = 100

func startWorker(wg *sync.WaitGroup, counter Counter) {
	defer wg.Done()
	for i := 0; i < iterationsNum; i++ {
		counter.Add()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	mutStr := newMutexStruct()

	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go startWorker(wg, mutStr)
	}

	wg.Wait()
	fmt.Printf("number of workers = %d, number of iterations = %d, counter = %d\n", workersNum, iterationsNum, mutStr.Get())

	wg = &sync.WaitGroup{}
	atStr := newMutexStruct()

	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go startWorker(wg, atStr)
	}

	wg.Wait()
	fmt.Printf("number of workers = %d, number of iterations = %d, counter = %d\n", workersNum, iterationsNum, atStr.Get())
}
