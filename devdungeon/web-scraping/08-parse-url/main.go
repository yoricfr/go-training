package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	complexUrl := "https://www.example.com/path/to/?query=123&this=that#fragment"
	parseUrl, err := url.Parse(complexUrl)
	if err != nil {
		log.Fatal()
	}

	// Print out URL pieces
	fmt.Println("Scheme: ", parseUrl.Scheme)
	fmt.Println("Host: ", parseUrl.Host)
	fmt.Println("Path: ", parseUrl.Path)
	fmt.Println("RawQuery: ", parseUrl.RawQuery)
	fmt.Println("Fragment: ", parseUrl.Fragment)

	// Get the query key/values as a map
	var customUrl url.URL
	customUrl.Scheme = "https"
	customUrl.Host = "google.com"
	newQueryValues := customUrl.Query()
	newQueryValues.Set("key1", "val1")
	newQueryValues.Set("key2", "val2")
	customUrl.Fragment = "bookmarkLink"
	customUrl.RawQuery = newQueryValues.Encode()

	fmt.Println("\nCustom URL:")
	fmt.Println(customUrl.String())
}
