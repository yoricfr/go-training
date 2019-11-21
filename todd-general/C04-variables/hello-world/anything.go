package main

import "fmt"

func bar() {
	fmt.Println("I'm in bar")
}

func main() {
	fmt.Println("Hello everyone, weather is good, temperature is fine. I hope you're doing well")
	foo()
	fmt.Println("Another thing")
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, i%2)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}
