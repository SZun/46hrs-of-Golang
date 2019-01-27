package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {
	s := square{
		length: 4.454635,
	}
	c := circle{
		radius: 4.23456,
	}
	info(s)
	info(c)
}

func info(s shape) {
	fmt.Println(s.area())
}

func (s square) area() float64 {
	return s.length * s.length
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
