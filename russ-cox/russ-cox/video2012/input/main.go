package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func main() {
	do(func(lang Lang) {
		data, _ := xml.MarshalIndent(lang, "", "\t")
		fmt.Printf("%s\n", data)
	})
}

func do(f func(Lang)) {
	input, _ := os.Open("lang.jso")
	dec := json.NewDecoder(input)
	for {
		var lang Lang
		err := dec.Decode(&lang)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		f(lang)
	}
}
