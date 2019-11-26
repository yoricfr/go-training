package main

import "fmt"

type person struct {
	name string
}

func (p person) get() string {
	return p.name
}

func main() {
	p := person{name: "Bill"}
	changeMe(&p)
	//fmt.Println(p.get())
}

func changeMe(p *person) {
	p.name = "Joe"
	fmt.Println(p.get())
}
