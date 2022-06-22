package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]struct{})
	res := make([]string, 0)
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		res = append(res, v)
		m[v] = struct{}{}
	}

	fmt.Println(res)
}
