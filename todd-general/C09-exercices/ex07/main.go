package main

import "fmt"

func main() {
	i := 12
	if i >= 20 {
		fmt.Println("greater or equal to 20")
	} else if i >= 10 {
		fmt.Println("greater or equal to 10")
	} else {
		fmt.Println("lower than 10")
	}

	switch {
	case i >= 20:
		fmt.Println(">=20")
	case i >= 10:
		fmt.Println(">=10")
	default:
		fmt.Println("<10")
	}
}
