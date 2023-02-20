package main

import "sync"

//Issue:- If we don't call done enough times or calling done is less than the number of add, this created deadlocks which is too deadly for the application
// ---------Deadlock----------------
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}

//This shows->fatal error: all goroutines are asleep - deadlock!
