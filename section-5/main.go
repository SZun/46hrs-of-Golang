package main

import "fmt"

func main() {
	type myOwnType int
	var x myOwnType
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
