package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("false")
	case 2 == 4:
		fmt.Println("2 == 4")
	case 3 == 3:
		fmt.Println("3 == 3")
		fallthrough
	case 7 == 9:
		fmt.Println("7 == 9")
	case 4 == 4:
		fmt.Println("4 == 4")
	}
}
