package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var xi []int
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Numbers needed.")
		return
	}
	for _, v := range args {
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			//	log.Fatal(err)
		} else {
			xi = append(xi, int(n))
		}
	}
	sort.Ints(xi)
	fmt.Println(xi)
}
