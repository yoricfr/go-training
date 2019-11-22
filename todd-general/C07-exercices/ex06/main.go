package main

import "fmt"

const (
	y1 = 2020 + iota
	y2
	y3
	y4
)

func main() {
	fmt.Printf("%d %d %d %d\n", y1, y2, y3, y4)
}
