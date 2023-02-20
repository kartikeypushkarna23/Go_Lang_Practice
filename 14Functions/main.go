package main

import "fmt"

func main() {
	v1 := adder(5, 4)
	v2 := proadder(1, 2, 3, 4, 5, 6, 6, 7)
	v3, v4 := promul(3, 2, 1)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
}

func adder(valone int, valtwo int) int {
	return valone + valtwo
}

func proadder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

//we can return more than one value and of any type
func promul(values ...int) (int, string) {
	total := 1
	for _, val := range values {
		total *= val
	}
	return total, "Hello from function"
}
