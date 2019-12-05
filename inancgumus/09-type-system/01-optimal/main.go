package main

import (
	"fmt"
	"math"
)

func main() {
	// an english letter (search web for: ascii control code)
	var char byte = 'A'
	fmt.Printf("%c\n", char)

	// a non-english letter (search web for: unicode codepoint)
	var unicode rune = 'ち'
	fmt.Printf("%c\n", unicode)

	// a year in 4 digits like 2040
	var year uint16 = 2040
	fmt.Println(year)

	// a month in 2 digits: 1 to 12
	var month uint8 = 12
	fmt.Println(month)

	// the speed of the light
	var lightSpeed uint32 = 670616629
	fmt.Println(lightSpeed)

	// angle of a circle
	var angle float32 = math.Pi / 3
	fmt.Println(angle)

	// PI number: 3.141592653589793
	var myPi float64 = 3.141592653589793
	fmt.Println(myPi)

}
