package advanced

import (
	"fmt"
	"time"
)



type ticketRequest struct {
	personID int
	numTickets int
	cost int
}

//simulate processing of ticket requests
func ticketProcessor(requests <- chan ticketRequest, results chan<- int){
	for req :=range requests {
		fmt.Printf("Processing %d ticket(s) of personId %d with total cost %d\n", req.numTickets, req.personID, req.cost)
		//simulate processing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}


func main() {
	numRequests :=5
	price:=5
	ticketRequests :=make(chan ticketRequest, numRequests)
	ticketResults :=make(chan int)
	// start ticket processor/worker
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	//send ticket requests
	for i:=range numRequests {
		ticketRequests <- ticketRequest{personID: i+1, numTickets:(i+1)*2, cost: (i+1)*price}
	}
	close(ticketRequests)

	for range numRequests {
		result := <-ticketResults
		fmt.Printf("Processed request for personId: %d\n", result)
	}
}







// ============= BASIC WORKER POOL PATTERN
// func worker(id int, tasks <-chan int , results chan<- int /*send-only channel*/) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		//Simulate work
// 		time.Sleep(time.Second)
// 		results <- task *2
// 	}
// }

// func main() {
// 	numWorkers :=3
// 	numJobs :=10
// 	tasks := make(chan int, numJobs) //buffered channel
// 	results :=make(chan int, numJobs)

// 	//create workers
// 	for i:= range numWorkers {
// 		go worker(i,tasks, results)
// 	}

// 	//Send values to the tasks channel
// 	for i := range numJobs {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	for range numJobs {
// 		result := <- results
// 		fmt.Printf("Result: %d\n", result)
// 	}

// }
