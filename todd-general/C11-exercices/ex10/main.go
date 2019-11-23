package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["coco_chanel"] = []string{"Pomme Depin", "forest", "trees"}
	delete(m, "moneypenny_miss")
	for k, v := range m {
		fmt.Println(k)
		for i, val := range v {
			fmt.Printf("\t\t%d\t%s\n", i, val)
		}
	}

}
