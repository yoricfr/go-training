package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Pick a number")
		return
	}
	n, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("%q is not a number. %v\n", args[1], err)
		return
	}
	if n%2 == 0 {
		msg := "\n"
		if n%8 == 0 {
			msg = " and it's divisible by 8\n"
		}
		fmt.Printf("%v is an even number"+msg, n)
	} else {
		fmt.Printf("%v is an odd number\n", n)
	}

}
