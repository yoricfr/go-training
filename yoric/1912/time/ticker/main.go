package main

import (
	"fmt"
	"time"
)

func main() {
	tc := time.NewTicker(time.Second)
	for {
		<-tc.C
		sec := time.Now().Second()
		fmt.Println(sec)
		if sec == 0 {
			tc.Stop()
			break
		}
	}
}
