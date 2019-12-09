package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func main() {
	debug.SetGCPercent(-1)
	report("initial memory usage")

	var a [size]int
	report("after declaring an array")

	b := a
	report("after copying the array")

	passArray(a)

	s1 := a[:]
	s2 := b[:1000]
	s3 := b[1000:10000]
	report("after slicings")

	passSlice(s1)

	fmt.Printf("Array's size : %d bytes.\n", unsafe.Sizeof(a))
	fmt.Printf("Array2's size: %d bytes.\n", unsafe.Sizeof(b))
	fmt.Printf("Slice1's size: %d bytes.\n", unsafe.Sizeof(s1))
	fmt.Printf("Slice2's size: %d bytes.\n", unsafe.Sizeof(s2))
	fmt.Printf("Slice3's size: %d bytes.\n", unsafe.Sizeof(s3))
}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
