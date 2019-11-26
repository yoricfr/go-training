package main

import "fmt"

func main() {
	s, i := bar()
	fmt.Println(foo(), s, i)
}

func foo() int {
	return 42
}

func bar() (string, int) {
	return "Hello", 33
}
