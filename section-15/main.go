package main

import "fmt"

func main() {
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f1 := foo(si...)
	fmt.Println(f1)
}

func foo(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
