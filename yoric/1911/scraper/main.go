package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
)

func main() {
	start := time.Now()
	response, err := http.Get("https://news.ycombinator.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	dataStrip := strip.StripTags(string(data))

	fmt.Println(dataStrip)
	fmt.Printf("%.2fs", time.Since(start).Seconds())
	// _, err = io.Copy(os.Stdout, response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
