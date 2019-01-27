package main

import "fmt"

func main() {
	f1 := func() int {
		return 3
	}
	fmt.Println(f1())
}
