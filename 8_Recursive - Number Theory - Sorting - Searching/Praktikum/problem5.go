package main

import "fmt"

func FindMinAndMax(arr []int) (min int, max int) {
	min = arr[0]
	max = arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {

	var arr = []int{5, 7, 4, -2, -1, 8}
	min, max := FindMinAndMax(arr)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

}
