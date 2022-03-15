package main

import "fmt"

func caesar(offset int, input string) string {
	offset = offset % 26
	var output string
	for i := 0; i < len(input); i++ {
		x := input[i] + byte(offset)
		if x > 122 {
			output = output + string(x-26)
		} else {
			output = output + string(x)
		}
	}
	return output
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
}
