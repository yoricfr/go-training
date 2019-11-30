package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Printing a line")
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}
	_ = resp
	fmt.Println(time.Since(start))
}
