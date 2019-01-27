package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)
	s2 := bar()
	fmt.Printf("%T\n", s2)
	i := s2()
	fmt.Println(i)
}

func bar() func() int {
	return func() int {
		return 451
	}
}

func foo() string {
	s := "Hello World"
	return s
}
