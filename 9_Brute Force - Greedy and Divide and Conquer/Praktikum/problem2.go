package main

import "fmt"

func moneyCoins(money int) []int {
	a := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}

	x := []int{}
	low := 1
	high := len(a) - 1

	for low <= high {
		med := (low + high) / 2
		if a[med] <= money {
			low = med + 1
		} else {
			high = med - 1
		}
	}
	if low == len(a) || a[low] != money {
		x = a[:low]
	}
	return x
}

func main() {
	fmt.Println(moneyCoins(432))
}
