package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("go rt 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("go rt 2")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("exiting")
}
