package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()
	fmt.Fprintf(h, "This string is going to be converted")
	fmt.Printf("Hash=%#x\n", h.Sum32())
}
