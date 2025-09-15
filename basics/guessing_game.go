package basics

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	source:= rand.NewSource((time.Now().UnixNano()))
	random:=rand.New(source)

	//generate a random number between 1 and 100
	target:= random.Intn(100) +1

	//welcome
	fmt.Println("welcome to the guessing game")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)

		//check if the guess if correct
		if guess == target {
			fmt.Println("Congratulations!")
			break
		} else if (guess <target) {
			fmt.Println("Too Low!")
		} else {
			fmt.Println("Too High!")
		}
	}
}
