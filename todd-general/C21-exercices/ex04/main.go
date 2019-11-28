package main

import (
	"fmt"
	"sync"
)

func main() {
	inc := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			n := inc
			n++
			inc = n
			fmt.Println(inc)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End value:", inc)
}
