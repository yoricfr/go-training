package main

import "fmt"

func main() {
	const minsPerDay = 24 * 60
	const weekDays = 7
	var weeks int = 2
	fmt.Printf("There are %d minutes in %d weeks", minsPerDay*weekDays*weeks, weeks)
}
