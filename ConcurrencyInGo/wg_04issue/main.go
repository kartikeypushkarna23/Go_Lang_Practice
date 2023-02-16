package main

import (
	"fmt"
	"sync"
)

// Issue:= If we pass the value not the reference to the function it creates panic.We must pass the reference to the wg.
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go work(wg)
	wg.Wait()
}

func work(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Work is done")
}
