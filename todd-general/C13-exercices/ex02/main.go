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
	persons := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for i, v := range persons {
		fmt.Println(i)
		for _, flavor := range v.favIceCream {
			fmt.Print(flavor, " ")
		}
		fmt.Println()
	}
	// fmt.Println(persons["Doh"].first)
	// for i, v := range persons["Doh"].favIceCream {
	// 	fmt.Println(i, v)
	// }
	// fmt.Println(persons["Dapp"].first)
	// for i, v := range persons["Dapp"].favIceCream {
	// 	fmt.Println(i, v)
	// }
}
