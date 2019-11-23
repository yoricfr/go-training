package main

import "fmt"

func main() {
	switch {
	case 10 >= 20:
		fmt.Println(">=20")
	case 20 >= 10:
		fmt.Println(">=10")
	default:
		fmt.Println("<10")
	}
}
