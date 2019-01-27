package main

import "fmt"

func main() {
	f1 := foo()
	n, s := bar()
	fmt.Println(f1, n, s)
}

func foo() int {
	return 3
}

func bar() (int, string) {
	return 3, "hello"
}
