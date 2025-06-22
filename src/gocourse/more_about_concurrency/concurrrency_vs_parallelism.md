# Concurrency vs Parallelism

## Introduction

- Concurrency: The ability of a system to handle multple tasks simultaneously. It involves managing multiple tasks that are in progress at the same time but not necessariliy executed at the same instant.

- Parallelism: The simultaneous execution of multiple tasks, typically using multiple processors or cores, to improve performance by running operations at the same time.

- Parallelism is all about executing multiple  tasks simultaneously, typically on multiple cores or processors and this is a subset of concurrency. 


## Code:

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	concurrencyVsParallelism1()
	concurrencyVsParallelism2()

}


func heavyTask(id int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Task %d is starting..\n", id)
	for range 100_000_000 {

	}
	fmt.Printf("Tasks %d is finished at time %v\n", id, time.Now())
}

func concurrencyVsParallelism2(){
	numThreads := 4

	runtime.GOMAXPROCS(numThreads)
	var wg sync.WaitGroup

	for i := range numThreads{
		wg.Add(1)
		go heavyTask(i, &wg)
	}

	wg.Wait()
}



func printNumbers(){
	for i := range 5 {
		fmt.Println(i, ":", time.Now())
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters(){
	for _,letter := range "ABCDE"{
		fmt.Println(string(letter), ":",time.Now())
		time.Sleep(500 * time.Millisecond)
	}
}

func concurrencyVsParallelism1(){
	go printNumbers()
	go printLetters()

	time.Sleep(3 * time.Second)
}
```


**How parallelism is implemented in GO ?**

- It's the go runtime. Go's runtimes scheduler can execute Go routines in parallel, taking advantage of multiple core processors.

- We can have processes that are executed concurrently without being parallel. And that happens when we have a single core CPU with time slicing. The single core CPU will divide time using time slicing and work on those multiple tasks simultaneously by dividing time and giving time to different functions, different tasks in a shared way. eg: So maybe 200 milliseconds to a tasks and then next 200 ms to another tasks and next 50 ms to the first task that it left earlier, and so on.


<kbd><img src="../../assets/course_materials/concurrency_vs_paralleism_2.png" alt="Concurrency vs Parallelism"/></kbd>

- **Practical Applications:**
    - Concurrency Use cases:
        - I/O bound tasks
        - Server Applications
    - Parallelism Use Cases
        - CPU Bound tasks
        - Scientific Computing

- **Challenges and Considerations :**
    - Concurrency Challenges
        - Synchronization: managing shared resources to prevent race conditions.
        - Deadlocks: avoid situations where tasks are stuck waiting for each other.
    - Parallelism Challenges
        - Data Sharing
        - Overhead
    - Performance Tuning
