package main

import (
	"fmt"
)

func deleteElem[T any](arr []T, i int) []T {
	return append(arr[:i], arr[i+1:]...)
}

func deleteElem1[T any](arr []T, i int) []T {
	return arr[:i+copy(arr[i:], arr[i+1:])]
}

func deleteElemNoOrder[T any](arr []T, i int) []T {
	arr[i] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return arr[:len(arr)-1]
}

func deleteElemGC[T any](arr []*T, i int) []*T {
	if i < len(arr)-1 {
		copy(arr[i:], arr[i+1:])
	}
	arr[len(arr)-1] = nil // or the zero value of T
	return arr[:len(arr)-1]
}

func deleteElemNoOrderGC[T any](arr []*T, i int) []*T {
	arr[i] = arr[len(arr)-1]
	arr[len(arr)-1] = nil
	return arr[:len(arr)-1]
}

func main() {
	// удаление элемента
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("      original slice: %v\n", a)
	a = deleteElem(a, 2)
	fmt.Printf("deleted 3-rd element: %v\n\n", a)

	// удаление элемента, другой способ
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("      original slice: %v\n", a)
	a = deleteElem1(a, 2)
	fmt.Printf("deleted 3-rd element: %v\n\n", a)

	// удаление элемента без сохранения порядка
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("      original slice: %v\n", a)
	a = deleteElemNoOrder(a, 2)
	fmt.Printf("deleted 3-rd element: %v (no order)\n\n", a)

	// удаление элемента из массива указателей (устанавливаем удаленный элемент в nil)
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := make([]*int, 10)
	for i := range b {
		b[i] = &a[i]
	}
	fmt.Printf("      original slice: %v\n", b)
	b = deleteElemGC(b, 2)
	fmt.Printf("deleted 3-rd element: %v\n\n", b)

	// удаление элемента из массива указателей (устанавливаем удаленный элемент в nil)
	// без сохранения порядка
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b = make([]*int, 10)
	for i := range b {
		b[i] = &a[i]
	}
	fmt.Printf("      original slice: %v\n", b)
	b = deleteElemNoOrderGC(b, 2)
	fmt.Printf("deleted 3-rd element: %v (no order)\n\n", b)
}
