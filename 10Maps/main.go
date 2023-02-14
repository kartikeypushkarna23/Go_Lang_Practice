package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)

	languages["Js"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PT"] = "Python"

	fmt.Println(languages)

	delete(languages, "RB")

	for key, value := range languages {
		fmt.Printf("For key %v , value is %v\n", key, value)
	}
}
