package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	go task1()
	go task2()
	go task3()
	go task4()
	time.Sleep(time.Second) //temporary solution
	fmt.Println("elapsed: ", time.Since(now))
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task1")
}

func task2() {
	fmt.Println("task2")
}

func task3() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task3")
}

func task4() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task4")
}
