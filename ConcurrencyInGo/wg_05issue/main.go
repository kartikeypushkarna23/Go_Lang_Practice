package main

import (
	"sync"
	"time"
)

// Issue:= Waitgroup Issue Panic
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Second)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}

//It Shows:= panic: sync: WaitGroup is reused before previous Wait has returned
