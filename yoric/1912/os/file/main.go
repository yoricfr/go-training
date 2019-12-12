package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// f, _ := os.Create("test.txt")
	// count := int(1.024e7)
	// for i := 0; i < count; i++ {
	// 	f.Write([]byte{byte('a' + i%24)})
	// }
	// f.Close()
	// fmt.Println("ok")

	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	count := int(1.024e7)
	for i := 0; i < count; i++ {
		//fmt.Fprintf(w, "%c", byte('a'+i%24))
		w.Write([]byte{byte('a' + i%24)})
	}
	err = w.Flush() // Don't forget to flush!
	if err != nil {
		log.Fatal(err)
	}

}
