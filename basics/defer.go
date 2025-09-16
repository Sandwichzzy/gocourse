package basics

import "fmt"


func main() {

	process(10)
}
// Use cases
// 1.resource cleanup
// 2.unlocking mutexes
// 3.logging and tracing
func process(i int){
	defer fmt.Println("Deffered i value:", i) //Deffered i value: 10 参数取决于延迟语句
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	fmt.Println("Normal execution statement")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i:", i) //Value of i: 11
}