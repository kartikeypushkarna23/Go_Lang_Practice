package main

import "fmt"

func main() {
	fmt.Println("Variable declaration in Go")
	var username string = "KP"
	fmt.Println(username)

	//implicit type
	var name = "KD"
	fmt.Println(name)
	fmt.Printf("Data type of name is:  %T\n", name)

	//no var style
	number := 30.023654772782 //This type of declaration we cannot use outside function
	fmt.Println(number)
	fmt.Printf("Data type of number is:  %T\n", number)

	//constant
	const Name string = "KS"
	fmt.Println(Name)

	name = "kartik"
	fmt.Println(name)
}
