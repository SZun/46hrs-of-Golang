package main

import "fmt"

type person struct {
	first string
	last  string
	ficf  []string
}

func main() {
	p1 := person{
		first: "sam",
		last:  "zun",
		ficf:  []string{"a", "b", "c"},
	}
	p2 := person{
		first: "jeff",
		last:  "zun",
		ficf:  []string{"a", "b", "c"},
	}
	fmt.Println(p1.first, p1.last)
	for _, v := range p1.ficf {
		fmt.Println(v)
	}
	fmt.Println(p2.first, p2.last)
	for _, v := range p2.ficf {
		fmt.Println(v)
	}
}
