package main

import (
	"encoding/xml"
	"fmt"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func main() {
	lang := Lang{"Go", 2009, "https://golang.org"}
	data, _ := xml.MarshalIndent(lang, "", "\t")
	fmt.Printf("%s\n", data)
}
