package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var number int64
	var bit, index int

	fl := flag.NewFlagSet("set-bit", flag.ContinueOnError)
	fl.Int64Var(&number, "n", 12345, "int64 number to set bit")
	fl.IntVar(&bit, "b", 1, "bit value (0 or 1)")
	fl.IntVar(&index, "i", 5, "index of bit to set (from 1 to 64)")

	if err := fl.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	if bit > 1 || bit < 0 || index > 64 {
		fl.Usage()
		return
	}

	// формируем битовую маску
	var mask int64 = 1 << (index - 1)
	if bit == 1 {
		fmt.Printf("           mask: %064b\n", mask)
		fmt.Printf("original number: %064b\n", number)
		// xor
		// 0^0=0
		// 0^1=1
		// 1^0=1
		// 1^1=0
		fmt.Printf("  result number: %064b\n", number^mask)
		return
	}

	fmt.Printf("           mask: %064b\n", mask)
	fmt.Printf("original number: %064b\n", number)
	// and not
	// 0&^0=0
	// 0&^1=0
	// 1&^0=1
	// 1&^1=0
	fmt.Printf("  result number: %064b\n", number&^mask)
}
