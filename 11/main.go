package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func intersetcion[T constraints.Ordered](arr1 []T, arr2 []T) []T {
	hash := make(map[T]int)
	for _, v := range arr1 {
		hash[v] += 1
	}
	for _, v := range arr2 {
		hash[v] += 1
	}

	res := make([]T, 0)
	for k, v := range hash {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	set1 := []int{1, 2, 3, 4, 5, 10}
	set2 := []int{6, 4, 3, 8}

	res := intersetcion(set1, set2)
	fmt.Println(res)
}
