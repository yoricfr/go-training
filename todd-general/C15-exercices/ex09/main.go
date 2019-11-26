package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	f := func() {
		fmt.Println(t)
	}
	tellMe("Joe", f)
}

func tellMe(name string, time func()) {
	fmt.Println("Hello", name)
	time()
}
