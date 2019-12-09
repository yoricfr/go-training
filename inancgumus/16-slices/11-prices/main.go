package main

import (
	"fmt"
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
	xheader := strings.Split(header, separator)
	for _, v := range xheader {
		fmt.Printf("%-12s", v)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 60))

	xdata := strings.Split(data, "\n")
	for _, e := range xdata {
		xline := strings.Split(e, ",")
		for _, v := range xline {
			fmt.Printf("%-12s", strings.Trim(v, " \n\t"))
		}
		fmt.Println()
	}

}
