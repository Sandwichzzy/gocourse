package basics

import "fmt"

// Practical Use Cases
// Setup Tasks
// Configuration
// Registering Component
// Database Initialization
// Best Practices
// Avoid Side Effects
// Initialization Order
// Documentation

func init() {
	fmt.Println("Initializing package1...")
}

func init() {
	fmt.Println("Initializing package2...")
}

func init() {
	fmt.Println("Initializing package3...")
}

func main() {

	fmt.Println("Inside the main function")

}