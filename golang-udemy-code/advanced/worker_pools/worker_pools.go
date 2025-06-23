package main

import (
	"fmt"
	"time"
)

func main() {
	workerMain()
	ticketSellingAndBuying()
}



// ----------- BASIC WORKER POOL PATTERN -------------
func worker(id int, tasks <- chan int, results chan <- int){
	for task := range tasks {
		fmt.Printf("Worker %d processing tasks %d\n", id, task)
		// Simulate work
		time.Sleep(time.Second)
		results <- task * 2
	}
	
}

func workerMain(){
	numWorkers := 3
	numJobs := 10

	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create workers
	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	// Send values to the tasks channel
	for i:= range numJobs{
		tasks <- i
	}
	close(tasks)

	// Collect the results
	for range numJobs{
		result, ok := <- results
		if !ok {
			fmt.Println("All values received. Closing the channel.")
			break
		}
		fmt.Println("Result:", result)
	}
}


// ------------- TICKET SELLING AND BUYING MECHANISM -------------
type ticketRequest struct {
	personID int
	numTickets int
	cost int
}

// simulate processing of ticket requests
func ticketProcessor(requests <- chan ticketRequest, results chan <- int){
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost %d\n", req.numTickets, req.personID, req.cost)
		// simulate processing time
		time.Sleep(500 * time.Millisecond)
		results <- req.personID
	}
}

func ticketSellingAndBuying(){
	numRequests := 5
	price := 5

	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	// start ticket processor / worker
	for range 3{
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send ticket requests
	for i:= range numRequests {
		ticketRequests <- ticketRequest{personID: i+1, numTickets: (i+1) * 2, cost: (i+1)*price}
	}
	close(ticketRequests)

	for range numRequests{
		fmt.Printf("Ticker for personID %d processed successfully!\n", <- ticketResults)
	}
}