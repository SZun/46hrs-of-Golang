package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	defer foo()
	if err != nil {
		// fmt.Println("err happened", err)
		// log.Println("err happened", err)
		// log.Fatalln(err)
		// log.Panicln(err)
		panic(err)
	}
}

func foo() {
	fmt.Println("foo")
}
