package main

import "fmt"

func main() {
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Rimma",
		last:  "Zun",
		age:   21,
	}
	fmt.Println(p2)
	fmt.Println(p2.first, p2.last, p2.age)
}
