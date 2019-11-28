package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	inc := 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			n := inc
			runtime.Gosched()
			n++
			inc = n
			fmt.Println(inc)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End value:", inc)
}
