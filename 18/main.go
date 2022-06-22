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

// счетчик с использование мьютексов
type mutexCounter struct {
	counter int64
	mutex   *sync.RWMutex
}

func newMutexStruct() *mutexCounter {
	return &mutexCounter{
		counter: 0,
		mutex:   &sync.RWMutex{},
	}
}

func (m *mutexCounter) Add() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.counter++
}

func (m *mutexCounter) Get() int64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.counter
}

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

var workersNum = 20
var iterationsNum = 100

// воркер увеличивает счетчик
func startWorker(wg *sync.WaitGroup, counter Counter) {
	defer wg.Done()
	for i := 0; i < iterationsNum; i++ {
		counter.Add()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	mutStr := newMutexStruct()
	wg.Add(workersNum)

	for i := 0; i < workersNum; i++ {
		go startWorker(wg, mutStr)
	}

	wg.Wait()
	fmt.Printf("number of workers = %d, number of iterations = %d, counter = %d\n", workersNum, iterationsNum, mutStr.Get())

	wg = &sync.WaitGroup{}
	atStr := newMutexStruct()
	wg.Add(workersNum)

	for i := 0; i < workersNum; i++ {
		go startWorker(wg, atStr)
	}

	wg.Wait()
	fmt.Printf("number of workers = %d, number of iterations = %d, counter = %d\n", workersNum, iterationsNum, atStr.Get())
}
