package main

import "fmt"

func main() {
	Days := []string{"Sun", "Tues", "Wed", "Fri", "Sat"}

	//I method
	for i := 0; i < len(Days); i++ {
		fmt.Println(Days[i])
	}

	//II method
	for j := range Days {
		fmt.Println(Days[j])
	}

	//III method
	for index, day := range Days {
		fmt.Printf("Index is %v & value is %v\n", index, day)
	}

	//IV method
	roguevalue := 1

	for roguevalue < 10 {
		fmt.Println(roguevalue)
		roguevalue++
	}

}
