package main

import "fmt"

func main() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	if radius >= 200 && isSphere {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}
