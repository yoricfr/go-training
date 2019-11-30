package main

import "fmt"

type Lang struct {
	Name string
	Year int
	URL  string
}

func main() {
	lang := Lang{"Go", 2009, "http://golang.org"}
	fmt.Printf("%+v\n", lang)
}
