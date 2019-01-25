package main

import "fmt"

func main() {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
	}
	fmt.Println(m)
	fmt.Println("3")
	v, ok := m["8"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["3"]; ok {
		fmt.Println(v)
	}
}
