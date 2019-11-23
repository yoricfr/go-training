package main

import "fmt"

func main() {
	if fmt.Print("Hello "); 5 > 3 {
		fmt.Println("Everyone")
	}
	if x := 4; x > 2 {
		x++
		fmt.Println(x)
	}
}
