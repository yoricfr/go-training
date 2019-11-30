package main

import "fmt"

func main() {
	myPrint("Hi, ", 42, "\n")
}

func myPrint(args ...interface{}) {

	// Using .(type)
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Printf("%d", arg)
		case string:
			fmt.Printf("%s", arg)
		}
	}

	// Using reflect
	// for _, arg := range args {
	// 	switch v := reflect.ValueOf(arg); v.Kind() {
	// 	case reflect.String:
	// 		os.Stdout.WriteString(v.String())
	// 	case reflect.Int:
	// 		os.Stdout.WriteString(strconv.FormatInt(v.Int(), 10))
	// 	}
	// }

}
