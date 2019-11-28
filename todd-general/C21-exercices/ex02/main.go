package main

import "fmt"

type person struct {
}

func (p *person) Speak() {
	fmt.Println("I speak")
}

type human interface {
	Speak()
}

func saySomething(h human) {
	h.Speak()
}

func main() {
	var p person
	saySomething(&p)
}
