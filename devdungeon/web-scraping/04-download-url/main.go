package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://www.devdungeon.com/archive")
	if err != nil {
		log.Fatal(err)
	}

	outFile, _ := os.Create("output.html")
	// f := os.NewFile(755, "test.txt")
	n, err := io.Copy(outFile, response.Body)
	_ = n
	_ = err
}
