package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	length := len(os.Args[1])
	fmt.Println(length)
	length2 := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length2)
}
