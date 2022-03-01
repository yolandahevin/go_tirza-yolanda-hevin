package main

import "fmt"

func playWithAsterisk(n int) {
	for i := 1; i <= n; i++ {
		spasi(n - i)
		bintang(i)
		fmt.Println()
	}
}

func spasi(n int) {
	for j := 1; j <= n; j++ {
		fmt.Print(" ")
	}
}

func bintang(n int) {
	for j := 1; j <= n; j++ {
		fmt.Print("* ")
	}

}

func main() {
	playWithAsterisk(5)
}
