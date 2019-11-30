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
	c := make(chan string)
	n := 0
	do(func(lang Lang) {
		n++
		go count(lang.Name, lang.URL, c)
	})

	timeout := time.After(time.Millisecond * 700)
	for i := 0; i < n; i++ {
		select {
		case msg := <-c:
			fmt.Print(msg)
		case <-timeout:
			fmt.Printf("Timeout\n")
			return
		}

	}

	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}

func count(name, url string, c chan<- string) {
	start := time.Now()
	b, _ := http.Get(url)
	n, _ := io.Copy(ioutil.Discard, b.Body)
	b.Body.Close()
	dt := time.Since(start).Seconds()
	c <- fmt.Sprintf("%-6s %9d [%.2fs]\n", name, n, dt)
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
