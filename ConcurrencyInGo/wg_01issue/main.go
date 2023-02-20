package main

import (
	"fmt"
	"sync"
	"time"
)

// Issue:- If we Don't call the method add before we call wait, wait will return immediately.
func main() {
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Go routine : Done")
	}()
	wg.Wait()
	fmt.Println("Executed Immediately")
}
