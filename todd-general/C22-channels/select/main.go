package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("exit")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			e <- i
		default:
			o <- i
		}
	}
	// close(e)
	// close(o)
	q <- 0
	// close(q)
}

func receive(e, o, q <-chan int) {
	// for v := range e {
	// 	println(v)
	// }
	// for v := range o {
	// 	println(v)
	// }
	// println(<-q)
	for {
		select {
		case n := <-e:
			fmt.Println("even:", n)
		case n := <-o:
			fmt.Println("odd:", n)
		case n := <-q:
			fmt.Println("quit:", n)
			return
		}
	}
}
