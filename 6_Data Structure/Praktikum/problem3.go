package main

import "fmt"

func PairSum(array []int, target int) []int {
	barisangka := map[int]int{}
	for i, a := range array {
		hitung := target - a
		if j, b := barisangka[hitung]; b {
			return []int{i, j}
		}
		barisangka[a] = i
	}
	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
}
