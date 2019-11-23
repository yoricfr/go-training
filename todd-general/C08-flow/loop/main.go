package main

import "fmt"

func main() {
	for i := 1; i < 100000; i++ {
		fmt.Printf("%020b\n", i)
	}
}
