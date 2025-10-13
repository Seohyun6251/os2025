package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n') // ignore error
	fmt.Println(i)
	// fmt.Println(err)
	log.Fatal(err) // report error
}
