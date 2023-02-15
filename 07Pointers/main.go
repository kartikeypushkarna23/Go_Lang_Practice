package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	// var ptr *int //default value of pointer int is nil
	// fmt.Println(ptr)

	mynum := 20
	var ptr = &mynum
	*ptr = *ptr * 2
	fmt.Println("New value is : ", mynum)

}
