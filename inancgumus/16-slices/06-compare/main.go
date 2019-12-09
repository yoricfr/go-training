package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	namesA := strings.Split("Da Vinci, Wozniak, Carmack", ", ")
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	sort.Strings(namesA)
	sort.Strings(namesB)
	if len(namesA) != len(namesB) {
		fmt.Println("They are not equal.")
		return
	}
	for i := range namesA {
		if namesA[i] != namesB[i] {
			fmt.Println("They are not equal.")
			return
		}
	}
	fmt.Println("They are equal.")
}
