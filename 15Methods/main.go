package main

import "fmt"

func main() {

	Kartik := User{"Kp", "kp@gmail.com", true, 18}
	fmt.Println(Kartik)
	fmt.Printf("Kartik details are: %v\n", Kartik)
	fmt.Printf("Kartik is %v & email is %v\n", Kartik.Name, Kartik.Email)

	Kartik.GetStatus()
	Kartik.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Printf("Is user active: %v\n", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Printf("Email of this user is: %v\n", u.Email)
}
