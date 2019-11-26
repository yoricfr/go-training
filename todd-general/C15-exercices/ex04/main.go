package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   40,
	}
	(&p1).talk()
	(&person{"Sponge", "Bob", 12}).talk()
}

func (p person) speak() {
	fmt.Println("Hello, I'm", p.first, p.last, "and I'm", p.age, "years old.")
}

func (p *person) talk() {
	fmt.Println(p.first, "is talking")
}
