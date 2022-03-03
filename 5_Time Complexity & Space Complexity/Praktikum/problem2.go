package main

import (
	"fmt"
	"math"
)

func pow(x, n float64) int {
	var test = math.Pow(x, n)
	return int(test)
}

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}
