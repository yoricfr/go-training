package main

import "fmt"

func main() {
	fmt.Println(multiplier(3)())
}

func multiplier(p int) func() int {
	return func() int { return 5 * p }
}
