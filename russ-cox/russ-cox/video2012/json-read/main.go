package main

import (
	"io"
	"os"
)

func main() {
	input, _ := os.Open("lang.json")
	size, _ := io.Copy(os.Stdout, input)
	_ = size
}
