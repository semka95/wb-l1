package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// BinarySearch is implementation of binary search algorithm, it returns
// index of item if it was found and -1 if item was not found, this is
// implementation from the book
func BinarySearch(list []int, item int) int {
	var low, mid, guess int
	high := len(list) - 1

	for low <= high {
		mid = (low + high) / 2
		guess = list[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// GBinarySearch is generic implementation of binary search algorithm,
// it returns index of item if it was found and -1 if item was not found
func GBinarySearch[T constraints.Ordered](list []T, item T) int {
	var low, mid int
	var guess T
	high := len(list) - 1

	for low <= high {
		mid = (low + high) / 2
		guess = list[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	myList := []int{1, 3, 5, 7, 9, 13, 15, 17, 19, 20, 25, 34, 44, 45, 55, 56, 67, 78, 90, 100}
	myFloatList := []float32{1.1, 1.6, 2.5, 4.5, 6.1, 7.6, 11.4, 15.5}
	myStringList := []string{"ab", "ac", "b", "bcd", "fav", "fbd"}

	fmt.Println(BinarySearch(myList, 9))
	fmt.Println(BinarySearch(myList, -1))

	fmt.Println(GBinarySearch(myList, 7))
	fmt.Println(GBinarySearch(myFloatList, 6.1))
	fmt.Println(GBinarySearch(myStringList, "fav"))
}
