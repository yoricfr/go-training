package main

import "fmt"

func main() {
	x := "Hello"
	fmt.Printf("%v %p\n", &x, &x)
}
