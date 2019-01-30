package main

import "fmt"

func main() {
	m := map[string][]string{
		"rimma": []string{`1`, `2`, `3`},
		"jeff":  []string{`1`, `2`, `3`},
		"sam":   []string{`1`, `2`, `3`},
	}
	m["deda"] = []string{`1`, `2`, `3`}
	delete(m, "deda")
	for i, v := range m {
		fmt.Println(i)
		for j, k := range v {
			fmt.Println(j, k)
		}
	}
}
