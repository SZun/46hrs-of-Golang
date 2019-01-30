package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	a = int(b)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
