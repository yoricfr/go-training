package main

import "fmt"

func main() {
	var (
		names     []string
		distances []int
		data      []byte
		ratio     []float64
		alives    []bool
	)
	fmt.Printf("%11s: %-10T len=%d nil?%t\n", "names", names, len(names), names == nil)
	fmt.Printf("%11s: %-10T len=%d nil?%t\n", "distances", names, len(names), names == nil)
	fmt.Printf("%11s: %-10T len=%d nil?%t\n", "data", data, len(data), data == nil)
	fmt.Printf("%11s: %-10T len=%d nil?%t\n", "ratio", ratio, len(ratio), ratio == nil)
	fmt.Printf("%11s: %-10T len=%d nil?%t\n", "alives", alives, len(alives), alives == nil)
	print(names, distances, data, ratio, alives)
}

func print(i ...interface{}) {
	for k, e := range i {
		fmt.Printf("%T len=%d nil?%t\n", e, 0, i[k] == nil)
	}
}
