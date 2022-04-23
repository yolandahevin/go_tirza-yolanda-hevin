package main

import "fmt"

func MaximumBuyProduct(money int, productPrice []int) {
	sum := 0
	i := 0
	for index, item := range productPrice {
		sum += item
		i = index
	}
	if money <= sum {
		fmt.Println(i)
	}
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
}
