package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Printf("initial: a = %d, b = %d\n", a, b)

	a, b = b, a
	fmt.Printf("swapped: a = %d, b = %d\n", a, b)
}
