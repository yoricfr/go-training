package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// DONT TOUCH THIS
	// 1h30m means: 1 hour 30 minutes
	t, _ := time.ParseDuration("1h30m")

	arg := os.Args[1]
	multi, _ := strconv.ParseFloat(arg, 64)
	t *= time.Duration(multi)

	// DONT TOUCH THIS
	fmt.Println(t)
}
