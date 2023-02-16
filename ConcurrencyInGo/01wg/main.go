package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup //creation
	wg.Add(3)             //Add

	go func() {
		defer wg.Done() //Calling Done
		fmt.Println("1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("2")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("3")
	}()

	wg.Wait() //calling wait
}
