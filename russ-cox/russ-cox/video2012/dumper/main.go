package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	h := hex.Dumper(os.Stdout)
	defer h.Close()
	//h.Write([]byte("Hello, get me the hex values of this string please."))
	fmt.Fprintf(h, "Hello, get me the hex values of this string please.")
}
