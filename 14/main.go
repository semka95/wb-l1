package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := []any{"hi", 42, func() {}, struct{}{}, true, 45.6}
	for _, v := range arr {
		v := reflect.ValueOf(v)
		fmt.Println(v.Kind().String())
	}
}
