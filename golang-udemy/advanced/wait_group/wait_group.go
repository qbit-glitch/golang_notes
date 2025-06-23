package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// workerWaitGroupMain()
	// workerWGChannelMain()
	PerformTaskMain()
}

// -------- BASIC EXAMPLE WITHOUT USING CHANNELS -------
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// wg.Add(1)     // -> WRONG PRACTICE
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("Worker %d finished\n", id)
}

func workerWaitGroupMain() {
	var wg sync.WaitGroup
	numWorkers := 3

	wg.Add(numWorkers) // THIS IS THE CORRECT WAY OF ADDING COUNTER TO WAIT GROUPS

	for i := range numWorkers {
		go worker(i, &wg)
	}

	wg.Wait() // blocking mechanism
	fmt.Println("All workers finished")
}

// ------------- EXAMPLE WITH CHANNELS ----------------

func workerWGChannel(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("WorkerID %d starting. \n", id)
	
	time.Sleep(time.Millisecond * 500) // simulate some work
	
	for task := range tasks {
		results <- task * 2
	}
	
	fmt.Printf("WorkerID %d finished.\n", id)
}

func workerWGChannelMain() {
	var wg sync.WaitGroup
	numWorkers := 3
	numJobs := 9
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	wg.Add(numWorkers)

	for i := range numWorkers {
		go workerWGChannel(i+1, tasks, results, &wg)
	}

	for i := range numJobs{
		tasks <- i+1
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Received:", result)
	}
}


//  ------- SIMULATING REAL-WORLD SCENARIO: CONSTURCTION -----------
type Worker struct{
	ID int
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Task)
}

func PerformTaskMain(){
	var wg sync.WaitGroup

	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks{
		worker := Worker{ID: i+1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}
	// Wait for all workers to finish
	wg.Wait()

	// Construction is finished
	fmt.Println("Construction is finished")
}
