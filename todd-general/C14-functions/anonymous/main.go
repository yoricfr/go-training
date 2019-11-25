package main

import "fmt"

func main() {

	func(n int) {
		fmt.Println("I am anonymous", n)
	}(3)
}
