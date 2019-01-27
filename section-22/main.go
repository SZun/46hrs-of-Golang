package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(eve, odd, quit)
	recieve(eve, odd, quit)
}

func recieve(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:\t", v)
		case v := <-o:
			fmt.Println("from the odd channel:\t", v)
		case v := <-q:
			fmt.Println("from the quit channel:\t", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e)
	// close(o)
	q <- 0
}
