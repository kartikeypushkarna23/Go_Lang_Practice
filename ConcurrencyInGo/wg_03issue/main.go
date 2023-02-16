package main

import "sync"

//Issue:- If we call done more times than add it going to results panic.
func main() {
	var wg sync.WaitGroup
	wg.Done()

}

//This shows:= panic: sync: negative WaitGroup counter
