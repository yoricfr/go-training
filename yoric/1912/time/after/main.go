package main

import (
	"fmt"
	"math"
	"time"
)

var c chan int

func handle(int) {
	fmt.Println("Entering function handle")
	var u float64
	for i := 0; i < 1e09; i++ {
		u = math.Sin(float64(i))
	}
	fmt.Println(u)
	fmt.Printf("Exiting function handle, u=%.10f\n", u)
}

func main() {
	select {
	case m := <-c:
		handle(m)
	case <-time.After(10 * time.Second):
		fmt.Println("timed out")
	}
}
