package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[10])
	fmt.Printf("%T\n", "Hello World"[1])
	fmt.Println(len("もしもし"))
	fmt.Println("もしもし"[10])
	fmt.Printf("%T\n", "もしもし"[0])

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98\x44"
	s := []byte(sample)

	fmt.Println(sample)
	for i, v := range sample {
		fmt.Printf("%d %x, ", i, v)
		//fmt.Print(sample[i], " ")
	}
	fmt.Println()
	for j := 0; j < len(sample); j++ {
		fmt.Printf("%x ", sample[j])
	}
	fmt.Println()
	fmt.Printf("% x\n", sample)
	fmt.Printf("%q\n", sample)
	fmt.Printf("%+q\n", sample)

	fmt.Println(s)
	for i := range s {
		fmt.Printf("%x ", sample[i])
		//fmt.Print(sample[i], " ")
	}
	fmt.Println()
	for j := 0; j < len(s); j++ {
		fmt.Printf("%x ", s[j])
	}
	fmt.Println()
	fmt.Printf("% x\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%+q\n", s)

	for j := 0; j < len(sample); j++ {
		fmt.Printf("%q ", sample[j])
	}
	fmt.Println()

	const placeOfInterest = `⌘`
	fmt.Printf("%s\n", placeOfInterest)

	fmt.Printf("%+q\n", placeOfInterest)

	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Println()

	const nihongo = "日本語"
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

	str := "Hello, 世界"
	for _, v := range str {
		r, s := utf8.DecodeRuneInString(string(v))
		fmt.Println(string(r), s)
	}
	for i, w := 0, 0; i < len(str); i += w {
		r, s := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c %d", r, s)
		w = s
	}

	hello := "Hello, world!"
	for i := 0; i < len(hello); i++ {

		fmt.Printf("%#U ", hello[i])
	}
	fmt.Println()
}
