package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	request, err := http.NewRequest("GET", "https://www.devdungeon.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer request.Body.Close()

	// Create new cookie
	myCookie := &http.Cookie{
		Name:  "cookieKey1",
		Value: "value1",
	}

	// Add the cookie to your request
	request.AddCookie(myCookie)

	fmt.Println(request.Cookies())
	fmt.Println(request.Header)

	// Do something with the request
	// client := &http.Client{}
	// client.Do(request)
}
