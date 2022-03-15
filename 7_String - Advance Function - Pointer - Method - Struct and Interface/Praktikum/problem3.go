package main

import "fmt"

func swap(a, b *int) {
	var tmp = *a
	*a = *b
	*b = tmp
}

func main() {
	a := 30
	b := 70

	swap(&a, &b)
	fmt.Println(a, b)
}
