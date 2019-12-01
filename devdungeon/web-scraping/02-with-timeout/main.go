package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{Timeout: 1 * time.Second}

	response, err := client.Get("http://www.free.fr")
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
