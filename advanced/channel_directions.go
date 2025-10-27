package advanced

import "fmt"

// Send-only channel: 只能用于向channel发送数据，不能从该channel接收数据。类型为 chan<- T。
// Receive-only channel: 只能用于从channel接收数据，不能向该channel发送数据。类型为 <-chan T。

// chan<- T：箭头指向 chan，表示数据流入channel（即发送到channel）。
// <-chan T：箭头从 chan 指出，表示数据从channel流出（即从channel接收）。

func main() {

	ch := make(chan int)
	producer(ch)
	consumer(ch)
}

// Send only channel
func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// Receive only channel
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received: ", value)
	}
}