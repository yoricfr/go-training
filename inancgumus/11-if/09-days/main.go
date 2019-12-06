package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const ()

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me a month name")
		return
	}
	m := strings.ToLower(args[1])
	mn := 0
	if m == "january" {
		mn = 1
	} else if m == "february" {
		mn = 2
	} else if m == "march" {
		mn = 3
	} else if m == "april" {
		mn = 4
	} else if m == "may" {
		mn = 5
	} else if m == "june" {
		mn = 6
	} else if m == "july" {
		mn = 7
	} else if m == "august" {
		mn = 8
	} else if m == "september" {
		mn = 9
	} else if m == "october" {
		mn = 10
	} else if m == "november" {
		mn = 11
	} else if m == "december" {
		mn = 12
	}
	if mn == 0 {
		fmt.Printf("%q is not a month\n", args[1])
		return
	}
	year := time.Now().Year()
	// fmt.Printf("Year:%d\n", year)
	leap := 0
	var nbj int
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		leap = 1
	}
	// fmt.Printf("Leap:%d\n", leap)
	if mn == 2 {
		nbj = 28 + leap
	} else {
		nbj = 31 - (mn-1)%7%2
	}
	fmt.Printf("%q has %d days.\n", args[1], nbj)
}
