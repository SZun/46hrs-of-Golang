package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

type hotdog int

func (s secretAgent) speak() {
	fmt.Println(s.last+",", s.first, s.last)
}

func (p person) speak() {
	fmt.Println(p.last+",", p.first, p.last)
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("person: ", h.(person).last)
	case secretAgent:
		fmt.Println("secretAgent: ", h)
	default:
		return
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Ms",
			last:  "MP",
		},
		ltk: true,
	}

	p1 := person{
		first: "Sam",
		last:  "Zun",
	}

	fmt.Println(p1)
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	bar(sa1)
	bar(p1)
	p1.speak()

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
