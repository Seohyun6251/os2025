package main

import "fmt"

func main() {

	subjects := [4]string{"Go", "javascript", "Python", "Linux"} // slice literal

	subjectsSlice := subjects[:3] //슬라이싱
	subjectsSlice[0] = "java"
	subjects[0] = "java"

	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("==================")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}

}
