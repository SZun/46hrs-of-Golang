package main

import "fmt"

func main() {
	foo()
	func() { fmt.Println("anon") }()
	func(x int) { fmt.Println(x) }(43)
}

func foo() {
	fmt.Println("foo ran")
}
