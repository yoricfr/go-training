package main

import (
	"fmt"
	"time"
)

func main() {
	digits := [12][5]string{
		{"███", "█ █", "█ █", "█ █", "███"}, // 0
		{"██ ", " █ ", " █ ", " █ ", "███"}, // 1
		{"███", "  █", "███", "█  ", "███"}, // 2
		{"███", "  █", "███", "  █", "███"}, // 3
		{"█ █", "█ █", "███", "  █", "  █"}, // 4
		{"███", "█  ", "███", "  █", "███"}, // 5
		{"███", "█  ", "███", "█ █", "███"}, // 6
		{"███", "  █", "  █", "  █", "  █"}, // 7
		{"███", "█ █", "███", "█ █", "███"}, // 8
		{"███", "█ █", "███", "  █", "███"}, // 9
		{"   ", " ▨ ", "   ", " ▨ ", "   "}, // :
		{"   ", "   ", "   ", "   ", "   "}, // empty
	}

	var hms [5]int
	tac := 2
	for {
		hms[0] = time.Now().Hour()
		hms[2] = time.Now().Minute()
		hms[4] = time.Now().Second()
		fmt.Printf("\033[2J")
		fmt.Printf("\033[93m")
		for i := range digits[0] {
			for j, v := range hms {
				if j == 1 || j == 3 {
					fmt.Print(digits[9+tac][i], "  ")
				} else {
					fmt.Print(digits[v/10][i], "  ")
					fmt.Print(digits[v%10][i], "  ")
				}
			}
			fmt.Println()
		}
		fmt.Printf("\033[0m")
		time.Sleep(time.Second)
		if time.Second%2 == 0 {
			tac = 3 - tac
		}
	}
}
