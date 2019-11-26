package main

import "fmt"

func main() {
	x := func() string {
		return "Yes"
	}
	fmt.Println(x())
}
