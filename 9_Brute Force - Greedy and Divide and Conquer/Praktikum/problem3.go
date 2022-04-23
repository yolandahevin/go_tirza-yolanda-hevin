package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	sort.Ints(knightHeight)
	sort.Ints(dragonHead)
	var total int
	for _, dragon := range dragonHead {
		l := len(knightHeight)
		if pos := sort.Search(l, func(i int) bool {
			return knightHeight[i] >= dragon
		}); pos < l {
			total += knightHeight[pos]
			knightHeight = knightHeight[pos+1:]
		}
	}
	fmt.Println(knightHeight)
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})
}
