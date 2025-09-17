package main

import "fmt"

func main() {

	// 1.Printing Functions
	// fmt.Print("Hello ")
	// fmt.Print("World!")
	// fmt.Print(12, 456)

	// fmt.Println("Hello ")
	// fmt.Println("World!")
	// fmt.Println(12, 456)

	// name := "John"
	// age := 25
	// fmt.Printf("Name: %s, Age: %d\n", name, age)
	// fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// 2.Formatting Functions
	// s := fmt.Sprint("Hello", "World!", 123, 456) //HelloWorld!123 456
	// fmt.Print(s)

	// s = fmt.Sprintln("Hello", "World!", 123, 456) //Hello World! 123 456
	// fmt.Print(s)
	// fmt.Print(s)

	// sf := fmt.Sprintf("Name: %s, Age %d", name, age)
	// fmt.Println(sf)
	// fmt.Println(sf)

	// 3.Scanning Functions
	// var name string
	// var age int

	// fmt.Print("Enter your name and age:")
	// // fmt.Scan(&name, &age)   //可以一个一个输入 不会停止执行
	// //fmt.Scanln(&name, &age) //一行内输入
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// 4.Error Formatting Functions

	err := checkAge(17)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}