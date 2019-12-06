package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Requires age")
		return
	}
	age, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil || age < 0 {
		fmt.Printf("wrong age: %q\n", args[1])
		return
	}
	if age > 17 {
		fmt.Println("R-Rated")
	} else if age > 13 {
		fmt.Println("PG-13")
	} else {
		fmt.Println("PG-Rated")
	}
}
