package main

import "fmt"

var name string
var age float64
var char = 'a'

func main() {

  clever := true
  var clash interface{}

	fmt.Printf("%T %T %T %T %T\n", name, age, clever, clash, char)
}
