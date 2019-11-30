package main

import (
	"fmt"
	"os"
)

// Fprint functions take a io.Writer as a first argument
// io.Writer is an interface that wrap the Write() method
// stdout is pointing to a file which is the standard output
func main() {
	st := "This is a string to be displayed through Fprint"
	fmt.Fprintln(os.Stdout, st)
}
