package main

import "fmt"

func main() {
	fmt.Println("Hello, how many bytes?")
	//fmt.Printf("%d bytes\n", n)
	s, e := abc(3, 8)
	fmt.Println(s, e)
}

func abc(p ...int) (int, error) {
	return p[0] * 2, nil
}
