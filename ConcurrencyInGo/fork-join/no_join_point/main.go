package main

import (
	"fmt"
	"time"
)

func main() {
	go work() //fork point
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done Waiting, main exits")
	//no join point
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing anything")
}
