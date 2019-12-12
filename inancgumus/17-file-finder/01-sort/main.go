package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]
	// Expecting ints
	// if len(args) == 0 {
	// 	fmt.Println("Numbers required")
	// }
	// nums := make([]int, len(args))
	// for i, v := range args {
	// 	nums[i], _ = strconv.Atoi(v)
	// }
	// sort.Ints(nums)
	// str := make([][]byte, len(nums))
	// for i, v := range nums {
	// 	s := strconv.Itoa(v)
	// 	str[i] = []byte(s)
	// }
	// _ = ioutil.WriteFile("out.txt", bytes.Join(str, []byte{' '}), 0755)

	// Expecting strings
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
	}
	for j := 0; j < 1e4; j++ {
		sort.Strings(args)
		// data := make([][]byte, len(args))
		// for i, v := range args {
		// 	data[i] = []byte(v)
		// }

		// err := ioutil.WriteFile("out.txt", bytes.Join(data, []byte{'\n'}), 0755)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		var data []byte
		for _, s := range args {
			data = append(data, s...)
			data = append(data, '\n')
		}

		err := ioutil.WriteFile("sorted.txt", data, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}
