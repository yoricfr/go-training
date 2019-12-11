package main

import (
	s "github.com/inancgumus/prettyslice"
)

func init() {
	s.PrintBacking = true // prints the backing array
	s.MaxPerLine = 10     // prints 10 slice elements per line
	s.Width = 60          // prints 60 character per line
}

func main() {
	names := make([]string, 5)
	s.Show("1st step", names)

	names = append(names, "einstein", "tesla", "aristo")
	s.Show("2nd step", names)

	names = make([]string, 0, 5)
	names = append(names, "einstein", "tesla", "aristo")
	s.Show("3d step", names)

	// #4
	moreNames := [...]string{"plato", "khayyam", "ptolemy"}
	names = append(names, moreNames[:2]...)
	s.Show("4th step", names)

	// #5
	clone := make([]string, 3, 5)
	copy(clone, names[len(names)-3:])
	s.Show("5th step (before append)", clone)
	clone = append(clone, names[:2]...)
	s.Show("5th step (after append)", clone)

	// #6
	sliced := clone[1:4:4]
	sliced = append(sliced, "hypatia")
	clone[2] = "elder"
	s.Show("6th step", clone, sliced)

}
