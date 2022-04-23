package main

import (
	"fmt"
	"sort"
)

func main() {
	go func() {
		text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labora et dolore magna aliqua"
		counter := make(map[string]int)
		for _, c := range text {
			if string(c) != " " {
				counter[string(c)]++
			}
		}
		fmt.Println(text)
		chars := make([]string, 0, len(counter))
		for chr := range counter {
			chars = append(chars, chr)
		}

		sort.Slice(chars, func(i, j int) bool {
			return counter[chars[i]] > counter[chars[j]]
		})
		for _, chr := range chars {
			fmt.Printf("%v: %v\n", chr, counter[chr])
		}
		var input string
		fmt.Scanln(&input)
	}()
	fmt.Scanln()
}
