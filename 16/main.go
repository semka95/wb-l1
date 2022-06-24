package main

import "fmt"

// реализация быстрой сортировки из грокаем алгоритмы
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	less := make([]int, 0, 1+len(arr)/2)
	greater := make([]int, 0, 1+len(arr)/2)

	for _, val := range arr[1:] {
		if val <= pivot {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	arr = append([]int{}, quickSort(less)...)
	arr = append(arr, pivot)
	arr = append(arr, quickSort(greater)...)
	return arr
}

func main() {
	fmt.Println(quickSort([]int{6, 4, 7, 12, 6, 34, 8, 90}))
}
