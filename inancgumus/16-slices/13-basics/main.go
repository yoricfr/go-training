package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "2 4 6 1 3 5"
	var xint []int

	for _, e := range strings.Fields(data) {
		if n, err := strconv.Atoi(e); err == nil {
			xint = append(xint, n)
		}
	}

	even := xint[:3]
	odd := xint[3:]

	fmt.Printf("%-13s : %v\n", "nums", xint)
	fmt.Printf("%-13s : %v\n", "evens", even)
	fmt.Printf("%-13s : %v\n", "odds", odd)
	fmt.Printf("%-13s : %v\n", "middle", xint[2:4])
	fmt.Printf("%-13s : %v\n", "first 2", xint[:2])
	fmt.Printf("%-13s : %v\n", "last 2", xint[len(xint)-2:])
	fmt.Printf("%-13s : %v\n", "evens last 1", even[len(even)-1:])
	fmt.Printf("%-13s : %v\n", "odds last 2", odd[len(odd)-2:])
}
