package main

import (
	"fmt"
	"strconv"
)

func main() {
	var foundMap = map[byte]int{}
	angka := "76523752"
	for i := 0; i < len(angka); i++ {
		character := angka[i]
		_, exist := foundMap[character]
		if !exist {
			foundMap[character] = 0
		}
		foundMap[character] += 1
	}
	for character, found := range foundMap {
		if found == 1 {
			angkaStr := fmt.Sprintf("%c", character)
			angka, _ := strconv.Atoi(angkaStr)
			fmt.Print("[", angka, "]")
			// fmt.Printf("%c\n", character)
		}
	}

}
