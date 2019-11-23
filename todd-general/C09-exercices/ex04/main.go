package main

import "fmt"

func main() {
	year := 1977
	for {
		fmt.Println(year)
		year++
		if year > 2019 {
			break
		}
	}
}
