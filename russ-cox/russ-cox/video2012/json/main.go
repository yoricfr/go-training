package main

import (
	"encoding/json"
	"fmt"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func main() {
	lang := Lang{"Go", 2009, "https://golang.org"}
	data, _ := json.Marshal(lang)
	fmt.Printf("%s\n", data)
}
