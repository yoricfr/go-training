package main

import "fmt"

func main() {
	c := make(chan bool)
	go func() { c <- true }()
	fmt.Println(<-c)
}
