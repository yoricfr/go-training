package main

import (
	"fmt"
	"os"
)

const usage = "Usage: [username] [password]"
const userName = "jack"
const password = "1888"
const u2 = "yoric"
const p2 = "blabla"

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	if args[1] != userName && args[1] != u2 {
		fmt.Printf("Access denier for %q\n", args[1])
		return
	}
	if (args[1] == userName && args[2] == password) || (args[1] == u2 && args[2] == p2) {
		fmt.Printf("Access granted to %q\n", args[1])
	} else {
		fmt.Printf("Invalid password for %q\n", args[1])
	}

}
