package advanced

import (
	"fmt"
	"sync"
	"time"
)

//ticketExample
type ticketRequest struct {
	personID int
	numTickets int
	cost int
}

//simulate processing of ticket requests
func (t *ticketRequest) ticketProcessor( wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Processing %d ticket(s) of personId %d with total cost %d\n", t.numTickets, t.personID, t.cost)
	//simulate processing time
	time.Sleep(time.Second)
	fmt.Printf("Processed request for personId: %d\n", t.personID)
}


func main() {
	var wg sync.WaitGroup
	numRequests :=5
	price:=5
	for i:=0; i< numRequests; i++ {
		wg.Add(1)
		request := ticketRequest{personID: i+1, numTickets:(i+1)*2, cost: (i+1)*price}
		go request.ticketProcessor(&wg)
	}

	wg.Wait()
	fmt.Println("All ticket requests processed.")
}

// CONSTRUCTION EXAMPLE
// type Worker struct {
// 	ID int
// 	Task string
// }

// // PerformTask simulates a worker performing a task
// func (w *Worker) PerformTask(wg *sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting task: %s\n", w.ID, w.Task)
// 	time.Sleep(time.Duration(w.ID)* time.Second) // Simulate time taken to perform the task
// 	fmt.Printf("Worker %d completed task: %s\n", w.ID, w.Task)
// }

// func main() {
// 	var wg sync.WaitGroup
// 		// Define tasks to be performed by workers
// 	tasks := [] string{"digging","laying brick","painting"}

// 	// Create and start workers
// 	for i,task := range tasks {
// 		worker := Worker{ID:i+1,Task:task}
// 		wg.Add(1)
// 		go worker.PerformTask(&wg)
// 	}
	
// 	// Wait for all workers to finish
// 	wg.Wait()
// 		// Construction is finished
// 	fmt.Println("All workers completed.")
// }


// EXAMPLE WITH CHANNELS
// func worker(id int,tasks <-chan int, results chan<- int, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Printf("WorkerID %d starting \n", id)
// 	time.Sleep(time.Second) //simulate some time spent on processing the task
// 	for task := range tasks {
// 		results <- task*2
// 	}
// 	fmt.Printf("WorkerID %d finished \n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers :=3
// 	numJobs :=5
// 	results :=make(chan int, numJobs)
// 	tasks := make(chan int, numJobs)

// 	wg.Add(numWorkers)

// 	for i:=range numWorkers{
// 		go worker(i+1,tasks,results, &wg)
// 	}

// 	for i:=range numJobs{
// 		tasks <- i+1
// 	}
// 	close(tasks)

// 	go func (){
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Println("Result received:", result)
// 	}
// }

// // ============ BASIC EXAMPLE WITHOUT USING CHANNELS
// func worker (id int, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	// wg.Add(1) // This should not be here; Add should be called before starting the goroutine !!
// 	fmt.Printf("Worker %d starting \n", id)
// 	time.Sleep(time.Second) //simulate some time spent on processing the task
// 	fmt.Printf("Worker %d finished \n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers :=3
// 	wg.Add(numWorkers) //this is correct way of adding counter to wait group !!
// 	//Launch workers
// 	for i:=1 ; i<=numWorkers; i++{
// 		go worker(i,&wg)
// 	}
// 	//Wait for all workers to finish
// 	wg.Wait() //blocks until the counter is zero
// 	fmt.Println("All workers completed.")
// }
