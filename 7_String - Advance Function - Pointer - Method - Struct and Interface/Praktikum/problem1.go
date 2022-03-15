package main

import (
	"fmt"
	"strings"
)

func compare(a, b string) string {
	var shortStr = a
	var longStr = b

	if strings.Contains(a, b) {
		strings.EqualFold(a, b)
		for i := 0; i < len(shortStr); i++ {
			if strings.Index(shortStr, longStr) != -1 {
				return b
			}
		}
	}
	return a
}

func main() {
	fmt.Println(compare("AKA", "AKASHI"))
}
