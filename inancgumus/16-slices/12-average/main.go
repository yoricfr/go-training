package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
	New York,150,3,2,200000
	Paris,200,4,3,400000
	Istanbul,500,10,5,1000000`

		separator = ","
	)
	for _, h := range strings.Split(header, ",") {
		fmt.Printf("%-13s", h)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 70))
	lines := strings.Split(data, "\n")
	var (
		loc [5]string
		sum [4]int
	)
	for i, line := range lines {
		fields := strings.Split(strings.TrimSpace(line), ",")
		loc[i] = fields[0]
		for j := 0; j < 4; j++ {
			n, _ := strconv.Atoi(fields[j+1])
			sum[j] += n
		}
		for k := 0; k < 5; k++ {
			fmt.Printf("%-13v", fields[k])
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("%-13s", "Average")
	for j := 0; j < 4; j++ {
		fmt.Printf("%-13.2f", float64(sum[j])/4.)
	}
	fmt.Println()
}
