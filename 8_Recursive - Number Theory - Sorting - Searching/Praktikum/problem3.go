package main

import "fmt"

func primaSegiEmpat(high, wide, start int) {
	finish := high * wide
	var y int
	var a int
	var g int
	for y = start; y <= finish; y++ {
		a = 0
		for g = 1; g <= y; g++ {
			if y%g == 0 {
				a++
			}
		}
		if a == 2 {
			fmt.Print(y, " ")
		}
		if y%100 == 0 {
			fmt.Printf("\t \n")
		}
	}
}

func main() {
	primaSegiEmpat(5, 2, 1)
}
