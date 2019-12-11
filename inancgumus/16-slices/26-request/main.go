package main

import (
	"fmt"
	"strings"
)

func main() {
	// There are 3 request totals per day (8-hourly)
	const N = 3

	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	reqs := []int{
		500, 600, 250, // 1st day: 1350 requests
		200, 400, 50, // 2nd day: 650 requests
		900, 800, 600, // 3rd day: 2300 requests
		750, 250, 100, // 4th day: 1100 requests
		// grand total: 5400 requests
	}

	// ================================================
	// #1: Make a new slice with the exact size needed.

	// _ = reqs // remove this when you start

	size := len(reqs) / N // you need to find the size.
	if len(reqs)%N != 0 {
		size++
	}
	daily := make([][]int, 0, size)

	// ================================================

	// ================================================
	// #2: Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]

	// _ = daily // remove this when you start
	for i := 0; i < len(reqs); i += N {
		daily = append(daily, reqs[i:i+N])
	}

	// ================================================
	// #3: Print the results

	// Print a header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Loop over the daily slice and its inner slices to find
	// the daily totals and the grand total.
	// ...
	var totalDay, totalGrand int
	for i, v := range daily {
		totalDay = 0
		for _, e := range v {
			fmt.Printf("%-10d %-10d\n", i+1, e)
			totalDay += e
		}
		fmt.Printf("%-10s %-10d\n", "   TOTAL: ", totalDay)
		totalGrand += totalDay
		fmt.Println()
	}
	fmt.Printf("%-10s %-10d\n", "   GRAND: ", totalGrand)
	// ================================================

}
