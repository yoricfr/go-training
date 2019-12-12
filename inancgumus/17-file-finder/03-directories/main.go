package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	paths := os.Args[1:]
	if len(paths) == 0 {
		fmt.Println("Paths Required")
		return
	}
	var data []byte
	for _, path := range paths {
		fd, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		names, err := fd.Readdirnames(0)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(path)
		data = append(data, path...)
		data = append(data, os.PathSeparator, '\n')

		for _, name := range names {
			// fmt.Printf("\t%s\n", name)
			data = append(data, '\t')
			data = append(data, name...)
			data = append(data, os.PathSeparator, '\n')
		}
		data = append(data, '\n')
	}
	err := ioutil.WriteFile("dirs.txt", data, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
