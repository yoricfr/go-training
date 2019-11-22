package main

import "fmt"

var n int = 3

type hotdog int

var h hotdog

func main() {
	h = 5
	fmt.Println(n, h)
	h = hotdog(n + 1)
	fmt.Println(h)
	fmt.Printf("%T - %T\n", n, h)
}
