package main

import "strings"
import . "fmt"
import "os"

func main() {
	Println(strings.ToLower(os.Args[1]))
}
