package main

import "fmt"

import "math"

func main() {
	radius := 10.
	var area float64

	area = math.Pi * radius * radius
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
}
