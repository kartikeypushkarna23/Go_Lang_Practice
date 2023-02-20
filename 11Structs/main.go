package main

import "fmt"

func main() {
	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}
	Kartik := User{"Kp", "kp@gmail.com", true, 18}
	fmt.Println(Kartik)
	fmt.Printf("Kartik details are: %v\n", Kartik)
	fmt.Printf("Kartik is %v & email is %v", Kartik.Name, Kartik.Email)
}
