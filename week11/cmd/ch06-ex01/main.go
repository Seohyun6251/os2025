package main

import "fmt"

func main() {

	subjects := []string{"Go", "javascript", "Python", "Linux"} // slice literal

	subjectsSlice := subjects[1:3] //슬라이싱
	//l = list()

	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("==================")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
