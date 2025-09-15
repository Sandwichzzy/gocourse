package main

import "fmt"


func main() {
// Switch statement in go is (switch case default) (fallthrough)
	// switch expression {
	// case value1:
	// Code to be executed if expression equals value1
	// fallthrough
	// case value2:
	// Code to be executed if expression equals value2
	// case value3:
	// Code to be executed if expression equals value3
	// default:
	// Code to be executed if expression does not match any value
	// }

	// Switch statement in other languages (switch case break default)
	// switch expression {
	// case value1:
	// Code to be executed if expression equals value1
	// break;
	// case value2:
	// Code to be executed if expression equals value2
	// break;
	// case value3:
	// Code to be executed if expression equals value3
	// break;
	// default:
	// Code to be executed if expression does not match any value
	// break;
	// }
		// fruit:="apple"

		// switch fruit{
		// case "apple":
		// 	fmt.Println("apple!")
		// case "banana":
		// 	fmt.Println("banana!")
		// default:
		// 	fmt.Println("unknown")
		// }

		//Multiple Conditions
		day := "Monday"

		switch day {
		case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
			fmt.Println("It's a weekday.")
		case "Sunday":
			fmt.Println("It's a weekend.")
		default:
			fmt.Println("Invalid day.")
		}

			// number := 15

			// switch {
			// case number < 10:
			// 	fmt.Println("Number is less than 10")
			// case number >= 10 && number < 20:
			// 	fmt.Println("Number is between 10 and 19")
			// default:
			// 	fmt.Println("Number is 20 or more")
			// }

			// num := 2

			// switch {
			// case num > 1:
			// 	fmt.Println("Greater than 1")
			// 	fallthrough
			// case num == 2:
			// 	fmt.Println("Number is 2")
			// default:
			// 	fmt.Println("Not Two")
			// }
			checkType(10)
			checkType(3.14)
			checkType("Heelo")
			checkType(true)
}

func checkType(x interface{}){
	switch x.(type){
	case int:
		fmt.Println("integer")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown type")
	}
}