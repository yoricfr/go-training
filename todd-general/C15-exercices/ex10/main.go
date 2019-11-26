package main

import "fmt"

func main() {
	m := myClose()
	for range "   " {
		fmt.Println(m())
	}
}

func myClose() func() int {
	var sum int
	return func() int {
		sum++
		return sum
	}
}
