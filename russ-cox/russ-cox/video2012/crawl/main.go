package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func main() {
	start := time.Now()
	do(func(lang Lang) {
		count(lang.Name, lang.URL)
	})
	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}

func count(name, url string) {
	start := time.Now()
	b, _ := http.Get(url)
	n, _ := io.Copy(ioutil.Discard, b.Body)
	b.Body.Close()
	fmt.Printf("%-6s %9d [%.2fs]\n", name, n, time.Since(start).Seconds())
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
