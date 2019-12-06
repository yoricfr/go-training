package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Give me a year number")
		return
	}
	y, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year\n", args[1])
		return
	}
	if y%4 != 0 {
		fmt.Printf("%d is not a leap year\n", y)
	} else if y%100 != 0 {
		fmt.Printf("%d is a leap year\n", y)
	} else if y%400 != 0 {
		fmt.Printf("%d is not a leap year\n", y)
	} else {
		fmt.Printf("%d is a leap year\n", y)
	}
}
