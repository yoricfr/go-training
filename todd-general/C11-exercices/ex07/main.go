package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	for i := range x {
		fmt.Println(x[i])
	}
	for i := range x {
		for j := range x[i] {
			fmt.Println(x[i][j])
		}
	}
}
