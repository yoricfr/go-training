package main

import "fmt"

func main() {
	n := 6.2
	fmt.Println(int(n))

	// n := int(6.99)
	// fmt.Println(int(6.99))

	var f float64 = 3
	fmt.Println(int(f))

	// error
	// x := int(3.1)
	// fmt.Println(x)

	// error
	// if int(3.1) > 3 {
	// 	fmt.Println("Weird")
	// }

	// Alias types
	var u8 uint8
	var by byte
	u8 = by
	_ = u8

	// int to float64
	fl := float64(int(4))
	_ = fl
	// float64 to int
	lf := int(float64(4))
	// lf = int(float64(4.01)) // Error
	_ = lf

	// ===> It means you can only convert from float64 to int when the fractional part is 0

	var sb []byte
	var s string
	// sb = s // Error
	sb = []byte(s) // OK
	_ = sb

	// Weird: no error
	orange := 65
	color := string(orange)
	fmt.Println(color)

	// color = ([]byte)(orange) // Error
	_ = string(sb)
	_ = []byte(s)

	fmt.Println(string([]byte{110, 111}))

	// Surprising
	fmt.Println(float64(3 / 2))

	var letter uint8 = 255
	fmt.Print(letter + 5)
}
