package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://www.devdungeon.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nSize of the web page: %.2f kb\n", float64(n)/1024)
}
