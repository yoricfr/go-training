package main

import "fmt"

import "os"

import "strconv"

import "math"

func main() {
	var radius, area float64

	radius, _ = strconv.ParseFloat(os.Args[1], 64)

	area = 4 * math.Pi * radius * radius

	// ADD YOUR CODE HERE
	// ...

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
}
