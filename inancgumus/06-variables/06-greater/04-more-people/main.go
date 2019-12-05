package main

import (
	"fmt"
	"os"
)

func main() {
	count := len(os.Args) - 1
	fmt.Printf("There are %d people!\n", count)
	p1 := os.Args[1]
	p2 := os.Args[2]
	p3 := os.Args[3]
	fmt.Printf("Hello great %s !\n", p1)
	fmt.Printf("Hello great %s !\n", p2)
	fmt.Printf("Hello great %s !\n", p3)
	fmt.Println("Nice to meet you all.")
}
