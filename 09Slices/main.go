package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go")

	var cars = []string{"wolkswagen", "tata", "toyota", "mahindra", "BMW"}
	fmt.Println(cars) //[wolkswagen tata toyota mahindra BMW]
	fmt.Println(len(cars))

	cars = append(cars[1:])
	fmt.Println(cars) //[tata toyota mahindra BMW]

	marks := make([]int, 4)
	marks[0] = 200
	marks[1] = 199
	marks[2] = 198
	marks[3] = 197

	fmt.Println(marks) //[200 199 198 197]
	fmt.Println(len(marks))
	marks = append(marks, 222, 252, 287, 299) //append will reallocate the memory initially slices is the size of 4
	fmt.Println(marks)                        //[200 199 198 197 222 252 287 299]
	fmt.Println(len(marks))

	sort.Ints(marks)
	fmt.Println(marks)

	//How to remove a value from slice based on index
	var bikes = []string{"jawa", "royalEnfield", "bajaj pulsar", "honda splender"}
	fmt.Println(bikes)
	var index int = 2
	bikes = append(bikes[:index], bikes[index+1:]...)
	//append also be used for the removing elements from slice as it reallocates the memory

}
