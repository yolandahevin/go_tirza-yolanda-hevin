package main

import (
	"fmt"
	"math"
)

func Frog(height []int) int {
	n := len(height)
	prev := 0
	prev2 := 0
	for i := 1; i < n; i++ {
		jOne := prev + int(math.Abs(float64(height[i])-float64(height[i-1])))
		if i > 1 {
			j := prev2 + int(math.Abs(float64(height[i])-float64(height[i-2])))
			cur_i := math.Min(float64(jOne), float64(j))
			prev2 = prev
			prev = int(cur_i)
		}
	}
	return prev
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
}
