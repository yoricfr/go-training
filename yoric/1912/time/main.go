package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	ca := time.After(time.Second)
	<-ca
	fmt.Println(time.UnixDate)
	fmt.Println(strings.Repeat("-", 10))
	<-time.After(time.Second)
	fmt.Println("The end")
}
