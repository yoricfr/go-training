package main

import "fmt"

type person struct {
	first, last string
	favIceCream []string
}

func main() {
	p1 := person{
		first:       "John",
		last:        "Doh",
		favIceCream: []string{"strawberry", "lemon"},
	}
	p2 := person{
		first:       "Smith",
		last:        "Dapp",
		favIceCream: []string{"vanilla", "chocolate", "pear"},
	}
	p := []person{p1, p2}
	//flavor := []string{p1.favIceCream, p2.favIceCream}
	for _, v := range p {
		fmt.Println(v.first, v.last, v.favIceCream)
	}
}
