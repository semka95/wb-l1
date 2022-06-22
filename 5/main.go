package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// читает данные из канала и выводит их
func worker(wg *sync.WaitGroup, in <-chan int) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("got %d\n", v)
	}
	fmt.Println("exit from worker")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wg := &sync.WaitGroup{}

	var timeToWork time.Duration
	fl := flag.NewFlagSet("worker-pool", flag.ContinueOnError)
	fl.DurationVar(&timeToWork, "d", 5*time.Second, "program will run for this amount of seconds")

	if err := fl.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	input := make(chan int, 1)
	wg.Add(1)
	go worker(wg, input)

	// отправляет в канал случайные числа с определенной периодичностью
	// по истечении определенного времени завершается
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		timer := time.NewTimer(timeToWork)
		defer timer.Stop()

		for {
			select {
			case <-ticker.C:
				a := rand.Int()
				input <- a
				fmt.Printf("sent: %d\n", a)
			case <-timer.C:
				fmt.Println("time exceeded, closing program")
				close(input)
				return
			}
		}
	}()

	wg.Wait()
}
