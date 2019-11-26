package main

import "fmt"

func main() {
	fmt.Println("Factorial 4 =", factorial(4))
}

func factorial(n int) int {
	switch {
	case n <= 0:
		return 0
	case n == 1:
		return 1
	default:
		return n * factorial(n-1)
	}
}
