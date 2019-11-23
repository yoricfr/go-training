package main

import "fmt"

func main() {
	for i := 65; i < 65+26; i++ {
		fmt.Println(i)
		for range "   " {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
