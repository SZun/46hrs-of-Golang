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
	m["6"] = 6
	if v, ok := m["3"]; ok {
		fmt.Println(v)
	}
	for i, v := range m {
		fmt.Println(i, v)
	}
	xi := []int{1, 2, 3, 4, 5}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
