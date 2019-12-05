package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("There are", len(os.Args)-1, "names")
}
