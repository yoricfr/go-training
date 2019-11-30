package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func main() {
	h := crc32.NewIEEE()
	mw := io.MultiWriter(os.Stdout, h)
	fmt.Fprintln(mw, "Trying to write to multiple writers...")
	fmt.Printf("hash=%#x\n", h.Sum32())
}
