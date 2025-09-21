package main

import "fmt"


func main() {

		num:=424
		fmt.Printf("%05d\n",num) //00424  min5char

		message:="Hello"
		fmt.Printf("|%10s|\n",message) //  |     Hello| min10char
		fmt.Printf("|%-10s|\n",message) // |Hello     |

		
		message1 := "Hello \nWorld!"
		message2 := `Hello \nWorld!` //反引号 不让转义字符执行 Hello \nWorld!
		fmt.Println(message1)
		fmt.Println(message2)


	  // sqlQuery := `SELECT * FROM users WHERE age > 30`
}
