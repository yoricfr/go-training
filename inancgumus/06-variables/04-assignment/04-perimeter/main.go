package main

import "fmt"

func main() {
	var (
		perimeter     int
		width, height = 5, 6
	)
	perimeter = (width + height) * 2
	fmt.Println(perimeter)
}
