package main

import "fmt"

// -Practical Use Cases
// Stateful Functions
// Encapsulation
// Callbacks
// -Usefulness of Closure
// Encapsulation
// Flexibility
// Readability
// -Considerations
// Memory Usage
// Concurrency
// -Best Practices
// Limit Scope
// Avoid Overuse
func main() {

		// sequence:=adder()
		// fmt.Println(sequence())
		// fmt.Println(sequence())
		// fmt.Println(sequence())

		// sequence2:=adder()
		// fmt.Println(sequence2())
		substractor:=func() func(int) int{
			countdown:=99;
			return func(x int) int{
				countdown-=x
				return countdown
			}
		}()

		//using closure substractor
		fmt.Println(substractor(1))
		fmt.Println(substractor(2))
		fmt.Println(substractor(3))
		fmt.Println(substractor(4))
		fmt.Println(substractor(5))
}

func adder() func() int{
	i:=0
	fmt.Println("previous value of i:",i)
	return func() int{
		i++
		fmt.Println("added 1 to i")
		return i
	}
}