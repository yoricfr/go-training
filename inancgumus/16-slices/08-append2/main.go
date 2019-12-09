package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		pizza      []string
		departure  []time.Time
		graduation []int
		light      []bool
	)

	pizza = append(pizza, "Cheese", "meat", "mushrooms", "eggs")
	departure = append(departure, time.Now(), time.Now().Add(time.Hour*24))
	graduation = append(graduation, 2004, 2009, 2015)
	light = append(light, true, false, true)
	fmt.Printf("pizza: %v\n", pizza)
	fmt.Printf("departure: %v\n", departure)
	fmt.Printf("graduation: %v\n", graduation)
	fmt.Printf("light: %v\n", light)
}
