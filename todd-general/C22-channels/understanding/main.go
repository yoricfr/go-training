package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 5
	c <- 6
	// go func() { c <- 5 }()
	// go func() { fmt.Println(<-c) }()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n-------\n", c)
	// d := make(chan<- int)
	// d <- 4
	// fmt.Println(<-d)
}
