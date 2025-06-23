package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu1, mu2 sync.Mutex
	go func(){
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
	}()

	go func(){
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
		time.Sleep(time.Second)
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		mu2.Unlock()
		mu1.Unlock()
	}()
	// time.Sleep(3 * time.Second)
	// fmt.Println("Main function Completed")
	select {}
	
	/* CORRECT CODE AVOIDING DEADLOCKS
	One of the Soln: Follow the same lock order

	go func(){
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
	}()

	go func(){
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
	}()
		
	time.Sleep(3 * time.Second)
	fmt.Println("Main function Completed")
	// select {}
	*/
}

