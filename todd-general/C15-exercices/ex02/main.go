package main

import "fmt"

func main() {
	fmt.Println(bar([]int{1, 2, 3}))
}

func foo(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	return foo(xi...)
}
