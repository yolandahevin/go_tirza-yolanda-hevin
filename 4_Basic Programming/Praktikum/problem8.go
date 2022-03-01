package main

import "fmt"

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Print(i*j, " ")
		}
	}
}

func main() {
	cetakTablePerkalian(9)
}
