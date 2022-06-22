package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func setIntersetcion[T constraints.Ordered](arr1 []T, arr2 []T) []T {
	m := make(map[T]int)

	for _, v := range arr1 {
		m[v] += 1
	}
	for _, v := range arr2 {
		m[v] += 1
	}

	res := make([]T, 0)
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	set1 := []int{1, 2, 3, 4, 5, 10}
	set2 := []int{6, 4, 3, 8}
	fmt.Printf(" first set: %v\nsecond set: %v\n", set1, set2)

	res := setIntersetcion(set1, set2)
	fmt.Printf("intersection: %v\n", res)
}
