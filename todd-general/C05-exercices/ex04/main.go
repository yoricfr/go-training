package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	fmt.Printf("value of x: %d\ntype of x: %T\n", x, x)
	x = 42
	fmt.Printf("value of x: %d\n", x)
}
