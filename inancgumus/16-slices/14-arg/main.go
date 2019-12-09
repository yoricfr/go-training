package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	fmt.Printf("%q\n", ships)

	args := os.Args[1:]
	if len(args) == 0 || len(args) > 2 {
		fmt.Println("Provide the [starting] and [stopping] positions")
		return
	}

	start, _ := strconv.Atoi(args[0])
	end := len(ships)
	if len(args) > 1 {
		end, _ = strconv.Atoi(args[1])
	}

	if start < 0 || start > end {
		fmt.Println("Wrong positions")
		return
	}

	fmt.Println(ships[start:end])
}
