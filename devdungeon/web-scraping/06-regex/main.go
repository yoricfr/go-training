package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	response, err := http.Get("http://www.devdungeon.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	dataInBytes, err := ioutil.ReadAll(response.Body)

	// Regular expression to find comments
	re := regexp.MustCompile("<!--(.|\n)*?-->")
	comments := re.FindAllString(string(dataInBytes), -1)
	if comments == nil {
		fmt.Println("No matches.")
	} else {
		for _, comment := range comments {
			fmt.Println(comment)
		}
	}
}
