package main

import "fmt"

func main() {
	xs1 := []string{"AK - Alaska", "AL - Alabama", "AR - Arkansas"}
	xs2 := []string{"AK - Alaska", "AL - Alabama", "AR - Arkansas"}
	fmt.Println(xs1)
	fmt.Println(xs2)
	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)
	for _, xs := range xxs {
		for _, v := range xs {
			fmt.Println(v)
		}
	}
}
