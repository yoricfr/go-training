package main

import "fmt"

func main() {
	x := 5
	fmt.Printf("%d\t%08b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%08b\t%#x\n", y, y, y)
}
