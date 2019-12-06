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
	year, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year\n", args[1])
		return
	}

	// var leap bool
	leap := year%400 == 0 || (year%100 != 0 && year%4 == 0)
	// if year%400 == 0 {
	// 	leap = true
	// } else if year%100 == 0 {
	// 	leap = false
	// } else if year%4 == 0 {
	// 	leap = true
	// }

	if leap {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
