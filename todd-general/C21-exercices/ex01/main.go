package main

import (
	"fmt"
	"math"
	"sync"
)

var wg1, wg2 sync.WaitGroup

func main() {
	wg1.Add(10)
	wg2.Add(10)
	go sin()
	go cos()
	wg1.Wait()
	wg2.Wait()
}

func sin() {
	for i := range "          " {
		fmt.Println("--sin:--", math.Sin(float64(i)*.113))
		wg1.Done()
	}
}
func cos() {
	for i := range "          " {
		fmt.Println("**cos:**", math.Cos(float64(i)*.113))
		wg2.Done()
	}
}
