package main

import "fmt"

// -Why Close Channels?
// Signal Completion
// Prevent Resource Leaks
// -Best Practices for Closing Channels
// (1)Close Channels Only from the Sender
// (2)Avoid Closing Channels More Than Once
// (3)Avoid Closing Channels from Multiple Goroutines
// -Common Patterns for Closing Channels
// Pipeline Pattern
// Worker Pool Pattern
// -Debugging and Troubleshooting Channel Closures
// lIdentify Closing Channel Errors
// Use `sync.WaitGroup` for Coordination

// ===
func producer(ch chan <- int){
	for i:=range 5{
		ch<-i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int){
	for val:=range in{
		if val%2==0 {
			out<- val
		}
	}
	close(out)
}

func main(){
	ch1:=make(chan int)
	ch2:=make(chan int)
	go producer(ch1)	
	go filter(ch1,ch2)
	for val :=range ch2{
		fmt.Println(val)
	}
}

// === 4- Avoid Closing Channels More Than Once
// func main(){
// 	ch:=make(chan int)
// 	go func(){
// 		close(ch)
// 		// close(ch)
// 	}()
// 	time.Sleep(time.Second) //panic: close of closed channel
// }

// === 1-Simple closing channel example
// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		fmt.Println(val)
// 	}

// }

// // 2- RECEIVING FROM A CLOSED CHANNEL
// func main(){
// 	ch:=make(chan int)
// 	close(ch)

// 	val,ok:=<-ch
// 	if !ok {
// 		fmt.Println("channel is closed")
// 		return
// 	}
// 	fmt.Println(val) //channel is closed
// }

// //  3- RANGE OVER CLOSED CHANNEL

// func main(){
// 	ch := make(chan int)
// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val:=range ch{
// 		fmt.Println("received:",val)
// 	}
// // output:
// // received: 0
// // received: 1
// // received: 2
// // received: 3
// // received: 4
// }

