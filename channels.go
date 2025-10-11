package main

import (
	"fmt"
	"time"
)

//Channel 是 Go 语言并发编程的核心特性之一，用于在不同的 goroutine 之间进行通信和同步。
func main() {

//variable := make(chan type) '<-' operator
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString // blocking because it is continuously trying to receive values, it is ready to receive continuous flow of data.
		greeting <- "World"
		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()

	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End of program.")

}
