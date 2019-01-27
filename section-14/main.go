package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("func expression")
	}
	f()

	mj := func(x int) {
		fmt.Println(x)
	}
	mj(23)
}
