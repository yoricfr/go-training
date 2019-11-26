package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//w := io.Writer()
	text := `This is a piece of bytes
it can even last for several lines

Another paragraph;
This is the end, my "friend".`
	io.Copy(os.Stdout, strings.NewReader("Bada"))
	fmt.Fprintln(os.Stdout, "Boom")
	os.Stdout.WriteString("!\n")
	if err := ioutil.WriteFile("tata.txt", []byte(text), 0755); err != nil {
		fmt.Println(err)
	}
	io.WriteString(os.Stdout, "End of writting file.\n")
}
