package main

import "fmt"

func main() {
	// DONT TOUCH THIS:
	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	// DONT TOUCH THIS:
	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	celsius *= Celsius(celsiusDegree * Temperature(factor))
	fahr *= Fahrenheit(fahrDegree * Temperature(factor))

	// ----------------------------------------
	// DONT TOUCH THIS
	fmt.Println(celsius, fahr)

	// YOU MAY REMOVE THESE WHEN YOU'RE DONE
	// _, _, _ = celsiusDegree, fahrDegree, factor
}
