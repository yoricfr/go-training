package main

import (
	"fmt"
	"os"
)

func main() {
	l := len(os.Args)
	if l == 1 {
		p("Give me args")
	} else if l == 3 {
		p("There are two: \"hi there\"")
	} else if l == 6 {
		p("There are 5 arguments")
	}
}

func p(s string) {
	fmt.Println(s)
}
