package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Start\n")
	timeout := time.After(time.Second * 2)
	r := <-timeout
	fmt.Printf("timeout %v\n", r)
	fmt.Printf("The end\n")
}
