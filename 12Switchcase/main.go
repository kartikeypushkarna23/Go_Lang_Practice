package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dicevalue is 1")
	case 2:
		fmt.Println("dicevalue is 2")
	case 6:
		fmt.Println("dicevalue is 6")
	default:
		fmt.Println("dicevalue collapsed")
	}
}
