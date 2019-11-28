package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	inc := int64(0)
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&inc, 1)
			//runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&inc))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End value:", inc)
}
