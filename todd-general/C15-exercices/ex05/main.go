package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	c float64
}
type CIRCLE struct {
	r float64
}
type SHAPE interface {
	area() float64
}

func main() {
	a := SQUARE{3}
	b := CIRCLE{1.5}
	INFO(a)
	INFO(b)
	a.area()
	b.area()
}

func (x SQUARE) area() float64 {
	return x.c * x.c
}
func (x CIRCLE) area() float64 {
	return math.Pi * x.r * x.r
}

func INFO(s SHAPE) {
	switch x := s.(type) {

	default:
		fmt.Printf("%T: ", x)
	}
	println(s.area())

}
