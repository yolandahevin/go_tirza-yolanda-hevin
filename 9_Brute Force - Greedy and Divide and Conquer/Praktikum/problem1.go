package main

import (
	"fmt"
	"math"
)

func SimpleEquations(a, b, c int) {
	d := b ^ 2 - 4*a*c
	var z float64 = float64(d)
	var y float64 = float64(a)
	var v float64 = float64(b)

	if d == 0 {
		x := -v + (math.Sqrt(z))/2*y
		fmt.Println("1 solution", x)
	} else if d < 0 {
		fmt.Println("no solution")
	} else {
		x1 := -v - (math.Sqrt(z))/2*y
		x2 := -v + (math.Sqrt(z))/2*y
		fmt.Println("solution", x1, x2)
	}
}

func main() {
	SimpleEquations(1, 2, 3)
}
