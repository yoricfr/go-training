package main

import "fmt"

func toSliceInterface(p []int) []interface{} {
	si := make([]interface{}, len(p))
	for i, v := range p {
		si[i] = v
	}
	return si
}

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[:1])
	fmt.Println(x[2:3])
	fmt.Println(toSliceInterface(x)...)
}
