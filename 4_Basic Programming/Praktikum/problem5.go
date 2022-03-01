package main

import "fmt"

func palindrome(input string) bool {
	for i := 0; i < len(input); i++ {
		j := len(input) - i - 1
		if input[i] != input[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
}
