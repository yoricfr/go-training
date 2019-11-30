package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func main() {
	// input, _ := os.Open("lang.jso")
	// dec := json.NewDecoder(input)
	// for {
	// 	var lang Lang
	// 	err := dec.Decode(&lang)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%v\n", lang)
	// }

	var langs []Lang
	b, _ := ioutil.ReadFile("lang.json")
	err := json.Unmarshal(b, &langs)
	_ = err
	for _, lang := range langs {
		fmt.Printf("%v\n", lang)
	}
}
