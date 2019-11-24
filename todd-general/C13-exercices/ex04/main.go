package main

import "fmt"

func main() {
	p := struct {
		name string
		age  int
	}{
		name: "Bob",
		age:  38,
	}
	fmt.Println(p.name, p.age)
}
