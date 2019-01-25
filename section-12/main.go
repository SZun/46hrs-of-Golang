package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Samuel",
		last:  "Zun",
		age:   21,
	}
	p2 := person{
		first: "Rimma",
		last:  "Zun",
		age:   21,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
