package main

import "fmt"

func main() {
	fmt.Println("Array in Go")

	//Declaration
	var car [4]string
	car[0] = "honda"
	car[1] = "suzuki"
	car[3] = "hyundai"

	fmt.Println(car) //[honda suzuki  hyundai]
	fmt.Println(len(car))

	var fruits = [3]string{"mango", "peaches", "grapes"}
	fmt.Println(len(fruits))
	fmt.Println(fruits)
}
