package main

import "fmt"

func main() {
	//Printing functions=>
	fmt.Println("hello from kp")
	fmt.Printf("hello from kp")
	fmt.Print("hello from kp")
	f := fmt.Sprint("hello from kp")
	b := fmt.Sprintf("hello from kp")
	c := fmt.Sprintln("hello from kp")
	fmt.Print(f)
	fmt.Print(b)
	fmt.Print(c)

	//Scanning Functions=>
	var y, z, age int
	var name string

	fmt.Scan(&y)
	fmt.Printf("Scan: Y = %d\n", y)

	fmt.Scanf("%s", &name)
	fmt.Scanf("%d", &age)
	fmt.Printf("Scanf: Name = %s , Age = %d\n", name, age)
	fmt.Scanln(&z)
	fmt.Printf("Scanln: Z = %d\n", z)

}
