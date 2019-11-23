package main

import "fmt"

func main() {
	var a1 [5]int = [5]int{1, 2, 3, 4, 5}
	var a2 [5]int = [5]int{2, 4, 6, 8, 0}

	fmt.Printf("%p\n", &a1)
	fmt.Printf("%p\n", &a2)

	a1 = a2
	fmt.Println(a1)

	fmt.Printf("%p\n", &a1)
	fmt.Printf("%p\n", &a2)

	s1 := "sos"
	fmt.Printf("ptr s1: %p\n", &s1)
	var s2 *string = &s1
	fmt.Printf("ptr s2: %p %s\n", s2, *s2)
}
