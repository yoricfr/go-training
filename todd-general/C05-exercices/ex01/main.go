package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("x:", x, "- y:", y, "- z:", z)
	fmt.Printf("x: %d - y: %s - z: %t\n", x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
