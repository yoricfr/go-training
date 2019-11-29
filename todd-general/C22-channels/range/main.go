package main

import "fmt"

func main() {
	c := make(chan int)

	// sending data on the chanel
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// retrieving the values from the chanel
	for v := range c {
		fmt.Println(v)
	}
}
