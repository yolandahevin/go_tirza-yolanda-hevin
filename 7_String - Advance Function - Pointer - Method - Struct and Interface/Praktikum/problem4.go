package main

import "fmt"

func getMinMax(number ...*int) (min int, max int) {
	min = *(number[0])

	for i := 0; i < len(number); i++ {
		if max < *(number[i]) {
			max = *number[i]
		}
		if min > *number[i] {
			min = *number[i]
		}
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Print("a1 :")
	fmt.Scan(&a1)
	fmt.Print("a2 :")
	fmt.Scan(&a2)
	fmt.Print("a3 :")
	fmt.Scan(&a3)
	fmt.Print("a4 :")
	fmt.Scan(&a4)
	fmt.Print("a5 :")
	fmt.Scan(&a5)
	fmt.Print("a6 :")
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Min ", min)
	fmt.Println("Nilai Max ", max)

}
