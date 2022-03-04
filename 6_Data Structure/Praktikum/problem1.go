package main

import "fmt"

func main() {
	list1 := []string{"kazuya", "jin", "lee"}
	list2 := []string{"kazuya", "feng", "jin"}
	list3 := list1

	for index, _ := range list2 {
		list3 = append(list3, list2[index])
	}

	occured := map[string]bool{}
	result := []string{}
	for e := range list3 {
		if occured[list3[e]] != true {
			occured[list3[e]] = true
			result = append(result, list3[e])
		}
	}
	fmt.Println(result)

}
