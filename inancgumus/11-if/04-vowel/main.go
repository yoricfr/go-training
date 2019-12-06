package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		args = os.Args
		l    = len(args) - 1
	)

	if l == 0 {
		fmt.Printf("Give me a letter\n")
	} else {
		lArg := len(os.Args[1])
		if lArg > 1 {
			fmt.Printf("Give me a letter\n")
		} else {
			c := os.Args[1]
			if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" {
				fmt.Printf("%q is a vowel\n", c)
			} else if c == "y" || c == "w" {
				fmt.Printf("%q is sometimes a vowel. sometimes not.\n", c)
			} else {
				fmt.Printf("%q is a consonant.\n", c)
			}
		}
	}
}
