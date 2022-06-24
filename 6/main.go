package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type goroutines struct {
	wg *sync.WaitGroup
}

// горутина проверяет закрылся ли канал пытаясь вычитать из него данные
// если вернулось стандартное значение (не ок), то завершаем функцию
func (g *goroutines) channelClose(ch <-chan int) {
	defer g.wg.Done()
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("goroutine checking if channel closed finished")
			return
		}
		fmt.Printf(" got: %d\n", v)
	}
}

// range завершится сам, как только канал закроется
func (g *goroutines) channelCloseRange(ch <-chan int) {
	defer g.wg.Done()
	for v := range ch {
		fmt.Printf(" got: %d\n", v)
	}
	fmt.Println("goroutine ranging on channel finished")
}

// используем отдельный канал для завершения горутины
func (g *goroutines) channelCloseCancelChannel(ch <-chan int, stop <-chan struct{}) {
	defer g.wg.Done()
	for {
		select {
		case v := <-ch:
			fmt.Printf(" got: %d\n", v)
		case <-stop:
			fmt.Println("goroutine with cancel channel finished")
			return
		}
	}
}

// используем сигнал из контекста для завершения горутины
func (g *goroutines) channelCloseContext(ctx context.Context, ch <-chan int) {
	defer g.wg.Done()
	for {
		select {
		case v := <-ch:
			fmt.Printf(" got: %d\n", v)
		case <-ctx.Done():
			fmt.Println("close context goroutine finished")
			return
		}
	}
}

// используем сигнал из контекста для завершения горутины
// в данном случае контекст завершается по таймауту, так как выполняется долгая операция
func (g *goroutines) channelCloseContextTimeout(ctx context.Context, ch <-chan int) {
	defer g.wg.Done()
	for {
		select {
		case <-time.After(10 * time.Second): // long operation
			fmt.Printf(" got: %d\n", <-ch)
		case <-ctx.Done():
			fmt.Println("context timeout goroutine finished")
			return
		}
	}
}

// запись случайных чисел в канал с определенной периодичностью
// завершается по сигналу контекста
func (g *goroutines) producer(ctx context.Context, ch chan int) {
	defer g.wg.Done()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			a := rand.Int()
			ch <- a
			fmt.Printf("sent: %d\n", a)
		case <-ctx.Done():
			fmt.Println("exiting from writer")
			ticker.Stop()
			return
		}
	}
}

func main() {
	// context timeout
	fmt.Println("Context timeout exceeded")
	ctx, cancel := context.WithCancel(context.Background())
	timeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ch := make(chan int, 1)
	g := goroutines{
		wg: &sync.WaitGroup{},
	}

	g.wg.Add(1)
	go g.producer(ctx, ch)
	go g.channelCloseContextTimeout(timeout, ch)

	g.wg.Wait()
	cancel()

	// canceling context
	fmt.Println("\nContext cancel")
	ctx, cancel = context.WithCancel(context.Background())
	ch = make(chan int, 1)

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.producer(ctx, ch)
	go g.channelCloseContext(ctx, ch)

	time.Sleep(3 * time.Second)
	cancel()
	close(ch)
	g.wg.Wait()

	// stopping channel
	fmt.Println("\nStop channel")
	ch = make(chan int, 1)
	stop := make(chan struct{})
	ctx, cancel = context.WithCancel(context.Background())

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.producer(ctx, ch)
	go g.channelCloseCancelChannel(ch, stop)

	time.Sleep(3 * time.Second)
	stop <- struct{}{}
	cancel()
	close(ch)
	g.wg.Wait()

	// ranging channel
	fmt.Println("\nRanging channel")
	ch = make(chan int, 1)
	ctx, cancel = context.WithCancel(context.Background())

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.producer(ctx, ch)
	go g.channelCloseRange(ch)

	time.Sleep(3 * time.Second)
	close(ch)
	cancel()
	g.wg.Wait()

	// closing channel
	fmt.Println("\nClosing channel")
	ch = make(chan int, 1)
	ctx, cancel = context.WithCancel(context.Background())

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.producer(ctx, ch)
	go g.channelClose(ch)

	time.Sleep(3 * time.Second)
	close(ch)
	cancel()
	g.wg.Wait()
}
