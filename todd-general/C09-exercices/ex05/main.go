package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%-4d%d\n", i, i%4)
	}
}
