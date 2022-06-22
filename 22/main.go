package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	a, ok := big.NewInt(0).SetString("5456513215645649854321321564878413215648649", 10)
	if !ok {
		fmt.Println("can't create first number")
		os.Exit(1)
	}
	b, ok := big.NewInt(0).SetString("4587132156798711316877984521256498789978441", 10)
	if !ok {
		fmt.Println("can't create second number")
		os.Exit(1)
	}

	fmt.Printf("1st number: %s\n2nd number: %s\n\n", a.String(), b.String())

	res := big.NewInt(0)
	fmt.Println("multiplex:", res.Mul(a, b))
	fmt.Println("division:", res.Div(a, b))
	fmt.Println("sum:", res.Add(a, b))
	fmt.Println("subtract:", res.Sub(a, b))
}
