package main

import "fmt"

func main() {
	cr := make(chan<- int)
	cs := make(<-chan int)

	fmt.Printf("%T, %T\n", cr, cs)
}
