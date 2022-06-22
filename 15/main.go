package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func createHugeString(size int) string {
	// используем буфер для эффективной конкатенации строк
	var b strings.Builder

	for i := 0; i < size; i++ {
		fmt.Fprint(&b, "界")
	}

	return b.String()
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// руна может занимать не один байт
	fmt.Println(utf8.RuneLen('界')) // 3

	// в данном случае мы срезаем по количеству байт, а не по количеству рун
	justString = v[:100]

	// преобразовываем строку в слайс рун
	r := []rune(v)
	// в даннам случае мы срезаем по количеству рун
	justString = string(r[:100])
}
func main() {
	someFunc()
}
