package advanced

import (
	"fmt"
	"time"
)

// - Why Use Buffered Channels?
// Asynchronous Communication
// Load Balancing
// Flow Control
// -Creating Buffered Channels
// make(chan Type, capacity)
// Buffer capacity 大的缓冲区减少阻塞的可能性 但增加内存使用量
// -Key Concepts of Channel Buffering
// Blocking Behavior
// Non-Blocking Operations
// Impact on Performance

// =========== BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY (first case)
// func main() {
// 	ch := make(chan int, 2)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		ch <- 2
// 	}()
// 	fmt.Println("Value: ", <-ch) //Blocks because the buffer is empty wait 2s
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("End of program.")
// }

// ================== BLOCKING ON SEND ONLY IF THE BUFFER IS FULL (second case)
func main() {
	// make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func() {
		fmt.Println("Goroutine 2 second timer started.")
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch) //ends <- starts
	}()
	// fmt.Println("Blocking starts")
	ch <- 3 // Blocks because the buffer is full
	// fmt.Println("Blocking ends")
	// fmt.Println("Received:", <-ch)
	// fmt.Println("Received:", <-ch)
}