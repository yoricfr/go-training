package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GoRoot", runtime.GOROOT())
	fmt.Println("Go OS", runtime.GOOS)
	fmt.Println("Go arch", runtime.GOARCH)
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(5))
	runtime.GC()
	fmt.Println("Compiler", runtime.Compiler)
}
