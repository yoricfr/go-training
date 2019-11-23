package main

import "fmt"

func main() {
	s := []int{3, 6, 8, 2, 4, 8, 4, 1}

	fmt.Printf("%v\t%p\n", s, &s[3])

	// Delete element at position 3
	s = append(s[:3], s[4:]...)
	fmt.Printf("%v\t\t%p\n", s, &s[3])
}
