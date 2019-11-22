package main

import "fmt"

var (
	x int    = 42
	y string = "James Bond"
	z bool   = true
)

func main() {
	s := fmt.Sprintf("x: %d, y: %s, z: %t\n", x, y, z)
	fmt.Print(s)
}
