package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i := range xi {
		fmt.Printf("%d ", xi[i])
	}
	fmt.Println()
	fmt.Printf("%T\n", xi)
}
