package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	arr := strings.Split(str, " ")

	var b strings.Builder
	b.WriteString(arr[len(arr)-1])
	for i := len(arr) - 2; i >= 0; i-- {
		b.WriteString(" ")
		b.WriteString(arr[i])

	}

	return b.String()
}

func main() {
	fmt.Println(reverseWords("snow dog sun"))
}
