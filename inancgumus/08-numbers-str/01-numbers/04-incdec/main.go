package main

import "fmt"

func main() {
	counter, factor := 45, 0.5

	counter++
	counter++
	counter++
	counter++
	counter++
	factor--
	factor--

	fmt.Printf("%g\n", float64(counter)*factor)
}
