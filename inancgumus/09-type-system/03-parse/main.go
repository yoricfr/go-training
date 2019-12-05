package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	val8, _ := strconv.ParseInt(os.Args[1], 10, 8)
	fmt.Println("int8 value is:", int8(val8))

	val16, _ := strconv.ParseInt(os.Args[2], 10, 16)
	fmt.Println("int16 value is:", int16(val16))

	val32, _ := strconv.ParseInt(os.Args[3], 10, 32)
	fmt.Println("int32 value is:", int32(val32))

	val64, _ := strconv.ParseInt(os.Args[4], 10, 64)
	fmt.Println("int64 value is:", val64)

	val8bit, _ := strconv.ParseInt(os.Args[5], 2, 8)
	fmt.Println(os.Args[5], "is:", int8(val8bit))
}
