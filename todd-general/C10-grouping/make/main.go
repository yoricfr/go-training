package main

import "fmt"

func main() {
	x := make([]int, 4, 5)
	x[0] = 2
	x[1] = 3
	x[2] = 5
	x[3] = 7

	fmt.Println(len(x), cap(x), x)
	fmt.Printf("array points to: %p\n", &x[0])

	x = append(x, 1)

	fmt.Println(len(x), cap(x), x)
	fmt.Printf("array points to: %p\n", &x[0])

	x = append(x, 2)

	fmt.Println(len(x), cap(x), x)
	fmt.Printf("array points to: %p\n", &x[0])
}
