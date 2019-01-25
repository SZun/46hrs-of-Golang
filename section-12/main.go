package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "Samuel",
		last:  "Zun",
	}
	p2 := person{
		first: "Rimma",
		last:  "Zun",
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}
