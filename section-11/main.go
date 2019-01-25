package main

import "fmt"

func main() {
	x := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:6])
	fmt.Println(x[1:6])
	fmt.Println(x)
}
