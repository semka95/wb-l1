package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var result int64
	wg := &sync.WaitGroup{}
	nums := []int64{2, 4, 6, 8, 10}

	for _, v := range nums {
		wg.Add(1)
		go func(num int64) {
			defer wg.Done()
			// потокобезопасная операция
			atomic.AddInt64(&result, num*num)
		}(v)
	}

	wg.Wait()

	fmt.Println(result)
}
