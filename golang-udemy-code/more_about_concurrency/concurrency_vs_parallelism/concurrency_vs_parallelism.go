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