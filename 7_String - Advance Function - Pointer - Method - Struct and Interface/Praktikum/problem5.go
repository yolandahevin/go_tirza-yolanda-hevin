package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var total int = 0
	for _, score := range s.score {
		total += score
	}
	var avg = float64(total) / float64(len(s.score))
	return avg
}

func (s Student) Max() (max int, name string) {
	max = 0
	indexmax := 0
	for i := 0; i < len(s.score); i++ {
		if max < s.score[i] {
			max = s.score[i]
		}
	}
	name = s.name[indexmax]
	return
}

func (s Student) Min() (min int, name string) {
	min = 100
	indexmin := 0
	for i := 0; i < len(s.score); i++ {
		if min > s.score[i] {
			min = s.score[i]
		}
	}
	name = s.name[indexmin]
	return
}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
