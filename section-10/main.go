package main

import "fmt"

func main() {
	jb := []string{"a", "b", "c", "d"}
	fmt.Println(jb)
	mp := []string{"e", "f", "g", "h"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
