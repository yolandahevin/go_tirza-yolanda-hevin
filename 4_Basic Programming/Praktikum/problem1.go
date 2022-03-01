package main

import "fmt"

func main() {
	var Lp, T, r float64
	const pi = 3.14
	fmt.Print("Masukkan tinggi tabung :")
	fmt.Scanln(&T)
	fmt.Print("Masukkan jari-jari tabung :")
	fmt.Scanln(&r)

	Lp = 2 * pi * r * (r + T)
	fmt.Println(Lp)

}
