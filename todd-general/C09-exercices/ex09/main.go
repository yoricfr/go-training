package main

import "fmt"

func main() {
	favSport := "ping-pong"
	var r string
	switch favSport {
	case "tennis":
		r = "I like it too"
	case "ping-pong":
		r = "My favorite sport!"
	case "natation":
		r = "what a great sport"
	default:
		r = "I don't know this sport"
	}
	fmt.Println(r)
}
