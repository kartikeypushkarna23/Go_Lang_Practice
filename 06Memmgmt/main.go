package main

import "fmt"

func main() {
	fmt.Println("Memory Management in Go")

	/*
		new()
		-Allocate memory but not initiaize
		-We will get a memory address
		-Zeroed storage(we cannot put data initially)
		---------------------------------------------
		make()
		-Allocate memory & init
		-we will get a memory address
		-non-zeroed storage

		$Garbage collection happens automatically
	*/
}
