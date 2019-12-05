package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hi %s\nHow are you?\n", os.Args[1])
}
