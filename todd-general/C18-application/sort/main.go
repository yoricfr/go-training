package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	xi := []int{2, 5, 38, 6, 9, 3, 66, 5, 4, 11, 26}
	s := "werdfiuw"
	xs := strings.SplitN(s, "", -1)
	//sort.Ints(xi)
	sort.IntSlice(xi).Sort()
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)
}
