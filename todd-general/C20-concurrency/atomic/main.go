package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	counter := int64(0)

	const gs = 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1.)
			fmt.Println("counter:", atomic.LoadInt64(&counter))
			//runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
