package main

import (
	"fmt"
	"sort"
)

func MaxSequence(arr []int) int {
	n := arr
	sort.Slice(n, func(i, j int) bool {
		return n[i] > n[j]
	})
	fmt.Println("sort   =>", n)

	var x []int = n[0:4]
	fmt.Println("slices =>", x)

	sum := 0
	for _, num := range x {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
}
