package main

import (
	"fmt"
	"sync"
)

type Counter interface {
	Add()
	Get() int64
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
