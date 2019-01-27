package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := foo(xi...)
	fmt.Println(x)
	foo()
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(i, v, sum)
	}
	fmt.Println(sum)
	return sum
}
