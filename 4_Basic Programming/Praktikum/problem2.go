package main

import "fmt"

func main() {
	var studentScore int
	var nama string

	fmt.Print("Masukan Nama Siswa : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Nilai Siswa ( Angka ) : ")
	fmt.Scan(&studentScore)

	if studentScore < 35 {
		fmt.Println("E")
	} else if studentScore < 50 {
		fmt.Println("D")
	} else if studentScore < 65 {
		fmt.Println("C")
	} else if studentScore < 80 {
		fmt.Println("B")
	} else if studentScore <= 100 {
		fmt.Println("A")
	} else if studentScore < 0 {
		fmt.Println("Nilai Invalid")
	} else if studentScore > 100 {
		fmt.Println("Nilai Invalid")
	} else {
		fmt.Println("Nilai Invalid")
	}

}
