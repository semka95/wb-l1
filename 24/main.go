package main

import (
	"fmt"
	"l1/24/point"
)

func main() {
	p1 := point.NewPoint(10.1, 15.2)
	p2 := point.NewPoint(-15.2, 3.1)

	line := point.Line{}
	fmt.Printf("distance between (10.1, 15.2) and (-15.2, 3.1) is %f\n", line.Distance(p1, p2))
}
