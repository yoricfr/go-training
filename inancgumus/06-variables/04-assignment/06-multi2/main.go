package main

import "fmt"

func main() {
	var (
		planet string  = "Mars"
		isTrue bool    = true
		temp   float64 = 19.5
	)

	planet, isTrue, temp = "Mercure", false, -23820
	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")
}
