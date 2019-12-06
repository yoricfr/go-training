package main

import "fmt"

func main() {
	age := 28

	if age > 60 {
		p("Getting older")
	} else if age > 30 {
		p("Getting wiser")
	} else if age > 20 {
		p("Adulthood")
	} else {
		p("Booting up")
	}
}

func p(s string) {
	fmt.Println(s)
}
