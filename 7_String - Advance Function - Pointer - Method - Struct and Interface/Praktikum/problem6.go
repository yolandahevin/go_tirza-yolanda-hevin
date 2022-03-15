package main

import (
	"fmt"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	s.score = s.score % 26

	for i := 0; i < len(s.name); i++ {
		x := s.name[i] + byte(s.score)
		if x > 122 {
			nameEncode = nameEncode + string(x-26)
		} else {
			nameEncode = nameEncode + string(x)
		}
	}
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	s.score = s.score % 26

	for i := 0; i < len(s.nameEncode); i++ {
		x := s.nameEncode[i] - byte(s.score)
		if x > 122 {
			nameDecode = nameDecode + string(x-26)
		} else {
			nameDecode = nameDecode + string(x)
		}
	}

	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	a.score = 3

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
