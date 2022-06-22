package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// воркер считывает число из канала и выводит его
func startWorker(ctx context.Context, wg *sync.WaitGroup, num int, in <-chan int) {
	defer wg.Done()
	var n int
	for {
		select {
		case n = <-in:
			fmt.Printf("worker № %d, got: %d\n", num, n)
		case <-ctx.Done():
			fmt.Printf("exiting from worker №%d\n", num)
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wg := &sync.WaitGroup{}

	var delay time.Duration
	var workersNum int
	fl := flag.NewFlagSet("worker-pool", flag.ContinueOnError)
	fl.DurationVar(&delay, "d", 1*time.Second, "message delay")
	fl.IntVar(&workersNum, "w", 2, "number of workers")

	if err := fl.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	input := make(chan int, 2)
	defer close(input)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(workersNum)

	// запускаем workersNum воркеров
	for i := 0; i < workersNum; i++ {
		go startWorker(ctx, wg, i+1, input)
	}

	// запись случайных чисел в канал с определенной периодичностью
	// завершается по сигналу контекста
	go func() {
		wg.Add(1)
		defer wg.Done()
		ticker := time.NewTicker(delay)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				a := rand.Int()
				input <- a
				fmt.Printf("writer sent: %d\n", a)
			case <-ctx.Done():
				fmt.Println("exiting from writer")
				return
			}
		}
	}()

	// по нажатию ctrl+c отправляем сигнал завершения всем горутинам через контекст
	// и ждем их завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("\ngot interrupt signal")
	cancel()
	wg.Wait()
}
