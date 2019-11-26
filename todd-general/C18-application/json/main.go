package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p := person{
		Age:   22,
		First: "Joe",
		Last:  "Star",
	}
	s, _ := json.Marshal(p)
	fmt.Println(string(s))
}
