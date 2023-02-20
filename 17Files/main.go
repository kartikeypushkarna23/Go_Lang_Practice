package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is first File."

	file1, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file1, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is: ", length)

	defer file1.Close()
	readfile("./myfile.txt")

}
func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text inside the file is: \n", databyte)
	fmt.Println("Text inside the file is :", string(databyte))
}
