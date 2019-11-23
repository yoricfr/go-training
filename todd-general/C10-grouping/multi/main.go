package main

import "fmt"

func main() {
	s1 := []string{"Jean", "Benichou", "developpeur"}
	s2 := []string{"Nicolas", "Fanfouay", "Menuisier"}
	fmt.Println(s1)
	fmt.Println(s2)
	s := [][]string{
		{"Jean", "Benichou", "developpeur"},
		{"Nicolas", "Fanfouay", "Menuisier"},
	}
	fmt.Println(s[0])
}
