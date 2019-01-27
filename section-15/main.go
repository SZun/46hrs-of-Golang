package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	sam := person{
		first: "Sam",
		last:  "Zun",
		age:   21,
	}
	sam.speak()
}

func (p person) speak() {
	fmt.Println("Hi my name is", p.first, p.last, "and I am", p.age, "years old.")
}
