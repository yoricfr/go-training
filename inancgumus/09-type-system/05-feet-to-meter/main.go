package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// ----------------------------
	// 1. Define Feet and Meters types below
	//    Their underlying type can be float64
	// ...
	type (
		Feet   float64
		Meters float64
	)

	// ----------------------------
	// 2. UNCOMMENT THE CODE BELOW THEN DON'T TOUCH IT
	var (
		feet   Feet
		meters Meters
	)

	// ----------------------------
	// 3. Get feet value from the command-line
	// 4. Convert it to an float64 first using ParseFloat
	// 5. Then, convert it into a Feet type
	// ... TYPE YOUR CODE HERE
	argFeet, _ := strconv.ParseFloat(os.Args[1], 64)
	feet = Feet(argFeet)

	// ----------------------------
	// 6. Uncomment the code below
	// 7. And, convert the expression to Meters type

	meters = Meters(feet) * 0.3048

	// ----------------------------
	// 8. UNCOMMENT THE CODE BELOW
	fmt.Printf("%g feet is %g meters.\n", feet, meters)

}
