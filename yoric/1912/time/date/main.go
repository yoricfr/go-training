package main

import (
	"fmt"
	"time"
)

func main() {
	y, m, d := time.Now().Date()
	fmt.Println(y, m, d)
}
