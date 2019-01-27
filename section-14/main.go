package main

import "fmt"

func main() {
	foo()
	bar("Sam")
	s1 := woo("MP")
	fmt.Println(s1)
	x, y := mouse("sam", "zun")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("foo")
}

func bar(s string) {
	fmt.Println("Hello, ", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello, from woo " + s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, " says hello")
	b := true
	return a, b
}
