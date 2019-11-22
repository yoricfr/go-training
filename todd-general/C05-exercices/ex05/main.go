package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Printf("value of x: %d\ntype of x: %T\n", x, x)
	x = 42
	fmt.Printf("value of x: %d\n", x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
