package main

import "fmt"

func main() {
	fmt.Println(func(n int) string {
		switch n {
		case 1:
			return "strawberry"
		case 2:
			return "apple"
		default:
			return "unknown"
		}
	}(1))
}
