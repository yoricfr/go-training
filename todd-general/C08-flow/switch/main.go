package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 2:
		i += 1
	case 3:
		i += 2
	case 4:
		i += 3
		fallthrough
	default:
		fmt.Println("Default")
	}
	fmt.Println(i)

	var t interface{} = "2"[0]
	fmt.Printf("%q %T\n", t, t)
}
