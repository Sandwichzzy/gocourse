package intermediate

import "fmt"

// -Practical Use Cases
// Mathematical Algorithms
// Tree and Graph Traversal
// Divide and Conquer Algorithms 分而治之
// -Benefits of Recursion
// Simplicity
// Clarity
// Flexibility
// -Considerations
// Performance
// Base Case
// -Best Practices
// Testing
// Optimization
// Recursive case
func main() {

	fmt.Println(factorial(5))
	fmt.Println(sumOfDigits(15))
	fmt.Println(sumOfDigits(15780))
}

func factorial(n int) int {
	//base case:factorial of 0 is 1
	if n==0 {
		return 1
	}
	//Recursive case : factorial of n is n*factorial(n-1)
	return n*factorial(n-1)
		// n * (n - 1) * (n-2) * factorial (n-3)..... factorial(0)
}

func sumOfDigits(n int) int {
	//base case
	if n<10 {
		return n
	}
	return n%10+sumOfDigits(n/10)
}