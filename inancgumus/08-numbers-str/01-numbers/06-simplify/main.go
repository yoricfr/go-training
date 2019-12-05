package main

import "fmt"

func main() {
	width, height := 10, 2

	width += 1
	width += height
	width--
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Println(width)
}
