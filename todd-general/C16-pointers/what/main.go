package main

import "fmt"

func main() {
	n := 12
	var p *int
	var q *(*int)

	p = &n
	q = &p
	fmt.Println(p, q)
}
