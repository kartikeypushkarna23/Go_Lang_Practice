package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input from user")

	//reference to pkg.go.dev & package name is -> bufio

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	//comma ok err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Your name is: ", input)
	fmt.Printf("Type of name is: %T", input)

}
