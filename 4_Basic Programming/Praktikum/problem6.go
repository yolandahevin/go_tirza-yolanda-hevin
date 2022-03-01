package main

import (
	"fmt"
	"math"
)

func pangkat(base, pangkat float64) int {
	var test = math.Pow(base, pangkat)
	return int(test)
}

func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}
