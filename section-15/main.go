package main

import "fmt"

func main() {
	f1 := func() int {
		return 451
	}
	f2 := foo(f1)
	fmt.Println(f2)
}

func foo(f func() int) int {
	return f()
}
