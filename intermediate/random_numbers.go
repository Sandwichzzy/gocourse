package intermediate

import (
	"fmt"
	"math/rand"
)

//伪随机数生成(PRNG):
//seed:生成随机数序列的起点 go 会autoseeding
//rand.Intn(n) Rand.Float64()
// 注意事项:
//确定性性质
//线程安全，对与并发不安全
//加密安全性


func main() {
	fmt.Println(rand.Intn(101)) //[0,101)
	// fmt.Println(rand.Intn(6) + 5) //[5,11)
	// // fmt.Println(val.Intn(101))
	// fmt.Println(rand.Float64()) // between [0.0,1.0)

	//  go 会autoseeding 不需要以下操作
	//  val:=rand.New(rand.NewSource(42)) //fix the seed
	//  val:=rand.New(rand.NewSource(time.Now().Unix()))


	for {
		// Show the menu
		fmt.Println("Welcome to the Dice Game!")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice (1 or 2): ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, please enter 1 or 2.")
			continue
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		die1:=rand.Intn(6)+1 //[1,6]
		die2:=rand.Intn(6)+1
		// show the results
		fmt.Printf("You rolled a %d and a %d.\n", die1, die2)
		fmt.Println("Total:", die1+die2)

		// Ask of the user wants to roll again
		fmt.Println("Do you want to roll again? (y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no.")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}

}
