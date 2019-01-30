package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	zun := person{
		first: "Sam",
		last:  "Zun",
	}
	fmt.Println(zun)
	changeMe(&zun, "Jeff")
	fmt.Println(zun)
	changeMe(&zun, "Rimma")
	fmt.Println(zun)
}

func changeMe(p *person, fn string) {
	p.first = fn
}
