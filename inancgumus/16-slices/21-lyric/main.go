package main

import (
	"fmt"
	"strings"
)

func main() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:
	// ...
	lyric = append([]string{"yesterday"}, lyric...)
	lyric = append(lyric[:8], append(lyric[13:], lyric[8:13]...)...)
	fmt.Printf("%s\n", lyric)
}
