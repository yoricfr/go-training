package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	x[0] = 2
	x[1] = 45
	x[2] = 6
	x[3] = 1
	x[4] = 9
	for i := range x {
		fmt.Printf("%02d ", x[i])
	}
	fmt.Println()
	fmt.Printf("%T\n", x)
}
