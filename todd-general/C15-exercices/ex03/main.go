package main

import "fmt"

func main() {
	foo()
}

func foo() {
	defer fmt.Println("Exit")
	fmt.Println("Start")
	for i := range []int{1, 2, 3, 4, 5} {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
