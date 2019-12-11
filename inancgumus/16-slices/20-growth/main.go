package main

import "fmt"

func main() {
	const size = 1e7

	var (
		xi      []int
		prevCap int
	)

	for i := 0; i < size; i++ {
		xi = append(xi, i)
		if prevCap != cap(xi) {
			fmt.Printf("Len = %10d \t cap = %10d \t growth = %.2f\n", len(xi), cap(xi), float64(cap(xi))/float64(prevCap))
			prevCap = cap(xi)
		}
	}
}
