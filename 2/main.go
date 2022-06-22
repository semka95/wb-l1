package main

import (
	"fmt"
	"sync"
)

// количество воркеров
var workersNum = 2

// воркер считывает число из канала и выводит его квадрат
func startWorker(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		fmt.Println(n * n)
	}
}

// вариант с использование воркеров
func concurrentSquareWorkers() {
	wg := &sync.WaitGroup{}
	wg.Add(workersNum)
	input := make(chan int, 2)

	// запускаем workersNum воркеров
	for i := 0; i < workersNum; i++ {
		go startWorker(input, wg)
	}

	// отправляем данные в канал
	nums := []int{2, 4, 6, 8, 10}
	for _, v := range nums {
		input <- v
	}

	// закрываем канал, и ждем завершения работы воркеров
	close(input)
	wg.Wait()
}

// вариант с запуском отдельной горутины на элемент
func concurrentSquare() {
	wg := &sync.WaitGroup{}
	nums := []int{2, 4, 6, 8, 10}

	for _, v := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(v)
	}

	wg.Wait()
}

func main() {
	concurrentSquare()

	concurrentSquareWorkers()
}
