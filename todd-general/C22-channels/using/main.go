package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	wg.Add(2)

	// sending to chanel
	go foo(c)

	// receiving from chanel
	go bar(c)

	wg.Wait()
	fmt.Println("End of program")
}

func foo(c chan<- int) {
	c <- 5
	wg.Done()
}

func bar(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}
