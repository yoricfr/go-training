package main

import "fmt"

type myint int

type person struct {
	age    int
	height float64
	name   string
}

func main() {
	var x myint = 3
	fmt.Println(x.double())
	p := person{
		age:    43,
		height: 1.8,
		name:   "Yoric",
	}
	fmt.Println(p)
}

func (i myint) double() int {
	return int(i) * 2
}
