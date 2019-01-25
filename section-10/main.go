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
	delete(m, "3")
	fmt.Println(m)
	if v, ok := m["4"]; ok {
		fmt.Println("value", v)
		fmt.Println(m["4"])
		delete(m, "4")
	}
	fmt.Println(m)
}
