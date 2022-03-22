package main

import "fmt"

func primeX(number int) int {
	for i := 1; i <= number; i++ {
		if number%i == 1 {
			i++
		}
	}
	if number%2 == 0 {
		return false
	}
	return true
}


}

func main() {
	fmt.Println(primeX(1))
}
