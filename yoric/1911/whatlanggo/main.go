package main

import (
	"fmt"
	"time"

	"github.com/abadojack/whatlanggo"
	"github.com/rylans/getlang"
)

func main() {
	sentences := []string{
		"El gato esta bien.",
		"Le petit chat est sur la table.",
		"This is my favorite color.",
		"Pwede mo ba akong tulungan? Sumama ka sa akin!",
		"mahal kita",
	}
	for _, sentence := range sentences {
		fmt.Printf("%s\n", sentence)
		method1(sentence)
		method2(sentence)
	}
}

func method1(s string) {
	fmt.Printf("Method1:\n")
	start := time.Now()
	info := whatlanggo.Detect(s)
	fmt.Println("Language:", info.Lang.String(), " Script:", whatlanggo.Scripts[info.Script], " Confidence", info.Confidence)
	fmt.Printf("%v\n---------\n", time.Since(start).Microseconds())
}

func method2(s string) {
	fmt.Printf("Method2:\n")
	start := time.Now()
	info := getlang.FromString(s)
	fmt.Println(info.LanguageCode(), info.Confidence())
	fmt.Printf("%v\n---------\n", time.Since(start).Microseconds())
}
