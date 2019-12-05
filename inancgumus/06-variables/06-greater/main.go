package main

import (
	"fmt"
	"os"
)

// main takes 3 arguments like this:
// go run main.go Hebra Jolly Bee
// and great them indicidually
func main() {
	// assign a new value to the string variable below
	var (
		name  = os.Args[1]
		name2 = os.Args[2]
		name3 = os.Args[3]
	)

	fmt.Println("Hello great", name, "!")
	fmt.Println("And hellooo to you magnificient", name2, "!")
	fmt.Println("Welcome", name3, "!")

	// changes the name, declares the age with 131
	name, age := "bilbo baggins", 131 // unknown age!

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("And, I love adventures!")

}
