package main

import "fmt"

func fibo(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibo(n-2) + fibo(n-1)
}

func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
}
