package main

import "fmt"

var x int

const (
	_  = iota
	kb = 1 << (10 * iota)
	mb
	gb
)

func main() {
	x = 1
	for x < 100 {
		fmt.Printf("%d %8b\n", x, x)
		x <<= 2
	}

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
