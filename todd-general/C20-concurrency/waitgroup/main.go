package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go foo()
	bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("foo ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("bar  ", i)
	}
}
