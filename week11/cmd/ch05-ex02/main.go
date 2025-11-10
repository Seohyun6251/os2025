// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("meetWeight.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	weeks := float64(len(numbers))
	fmt.Printf("평균 고기 소비량: %0.2f\n", sum/weeks)
}
