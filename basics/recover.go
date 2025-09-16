package basics

import "fmt"

// Practical Use Cases
// 1.Graceful Recovery
// 2.Cleanup
// 3.Logging and Reporting
// Best Practices
// 1.Use with Defer
// 2.Avoid Silent Recovery
// 3.Avoid Overuse

func main() {

	process()
	fmt.Println("Return from process")

}

func process(){
	defer func() {
		// if r:=recover();r!=nil{
		r:=recover()
		if r!=nil{
			fmt.Println("Recovered",r)
		}
	}()

	fmt.Println("startProcess")
	panic("something went wrong!")
	// fmt.Println("End process")
	
}