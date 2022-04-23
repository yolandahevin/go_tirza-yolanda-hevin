package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	is := true
	max := math.Sqrt(float64(n))
	for i := 2; i <= int(max); i++ {
		if n%i == 0 {
			is = false
			break
		}
	}
	return is
}

func primeX(number int) int {
	if number == 0 {
		return 0
	}
	count := 0
	i := 2
	for ; ; i++ {
		if isPrime(i) {
			count++
		}
		if count == number {
			break
		}
	}
	return i

}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
