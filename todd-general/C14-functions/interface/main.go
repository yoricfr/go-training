package main

import (
	"fmt"
	"strconv"
)

type paper struct {
	size  int
	Color string
}

type object interface {
	getWieght() float64
}

type myint int

func (n myint) addTwo() myint {
	return n + 2
}

func (p paper) getWieght() float64 {
	return float64(p.size) * 1.5
}

func getName(o object) string {
	return "I weight " + strconv.FormatFloat(o.getWieght(), 'g', 2, 64) + " pounds"
}

func main() {
	n := myint(10)
	fmt.Println(n.addTwo().addTwo().addTwo())
	p := paper{
		size:  4,
		Color: "orange",
	}
	fmt.Printf("%T\n", p)
	fmt.Println(p.getWieght())
	var o object
	o = p
	fmt.Printf("%T\n", o)
	// switch o.(type) {
	// case paper:
	// 	fmt.Println(o.(paper).color)
	// }
	fmt.Println(getName(o))
}
