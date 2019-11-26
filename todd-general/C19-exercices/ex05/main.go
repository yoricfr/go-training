package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type usersByAge []user
type usersByLast []user

func (xu usersByAge) Len() int {
	return len(xu)
}
func (xu usersByAge) Less(i, j int) bool {
	return xu[i].Age < xu[j].Age
}
func (xu usersByAge) Swap(i, j int) {
	xu[i], xu[j] = xu[j], xu[i]
}
func (xu usersByLast) Len() int {
	return len(xu)
}
func (xu usersByLast) Less(i, j int) bool {
	return xu[i].Last < xu[j].Last
}
func (xu usersByLast) Swap(i, j int) {
	xu[i], xu[j] = xu[j], xu[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println("==============\nInitial users data:")
	fmt.Println(users)

	// your code goes here
	fmt.Println("==============\nSorting by age:")
	ua := usersByAge(users)
	sort.Sort(ua)
	display(ua)
	fmt.Println("==============\nSorting by Last:")
	ul := usersByLast(users)
	sort.Sort(ul)
	display(ul)
}

func display(xu []user) {
	for i, v := range xu {
		fmt.Println("User number #", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, say := range v.Sayings {
			fmt.Println("\t\t", say)
		}
	}
}
