package main

import "fmt"

func main() {
	f1 := foo()
	fmt.Println(f1())
}

func foo() func() int {
	return func() int {
		return 451
	}
}
