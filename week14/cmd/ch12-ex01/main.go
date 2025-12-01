package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("1sr defer")
	defer fmt.Println("2nd defer")
	defer fmt.Println("3rd defer")
	fmt.Println("main logic")
}
