package main

import "fmt"

// записывает числа в канал из массива
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// читает числа из канала, возводит в квадрат и передает в другой канал
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	c := gen(2, 3, 5, 10, 45, 7, 5)
	out := sq(c)

	// выводим результат из выходного канала
	for v := range out {
		fmt.Println(v)
	}
}
