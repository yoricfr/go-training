package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	last := os.Args[2]
	fmt.Printf("Your name is %s and your last name is %s.\n", name, last)
}
