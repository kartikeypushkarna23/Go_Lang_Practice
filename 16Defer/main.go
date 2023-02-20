package main

import "fmt"

//defer keyword takes the function or statement & execute at the last when all the statements executed.
func main() {
	defer fmt.Println("Hello")
	fmt.Println("World")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
// Output:=
// World
// 4
// 3
// 2
// 1
// 0
// Hello