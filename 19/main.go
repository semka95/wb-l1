package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	// преобразовываем строку в слайс рун
	runes := []rune(str)
	var b strings.Builder

	for i := len(runes) - 1; i >= 0; i-- {
		b.WriteRune(runes[i])
	}

	return b.String()
}

func main() {
	fmt.Println(reverseString("日本語 test"))
}
