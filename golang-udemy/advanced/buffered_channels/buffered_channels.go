package main

import (
	"fmt"
	"time"
)

func main() {

	// blockingOnSendDemo()
	blockingOnReceiveDemo()


}

func blockingOnSendDemo(){
	// ===== BLOCKING ON SEND ONLY IF THE BUFFER IS FULL =======
	// variable = make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func(){
		time.Sleep(2 * time.Second)
		fmt.Println("Received:",<-ch)
	}()
	fmt.Println("Blocking Starts")
	
	ch <- 3			// Blocks because the buffer is full
	fmt.Println("Blocking Ends")
	
	fmt.Println("Received:",<-ch)
	fmt.Println("Received:", <-ch)

	fmt.Println("Buffered Channels")
}


func blockingOnReceiveDemo(){
	// ====== BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY =====
	ch := make(chan int, 2)
	go func ()  {
		time.Sleep(2 * time.Second)
		ch <- 1
		ch <- 2
		fmt.Println("Value sent to channel successful.")
		fmt.Println("Go Routine completed")

	}()

	fmt.Println("Value:", <- ch)
	fmt.Println("Value:", <- ch)

	fmt.Println("End of Blocking on Receive Demo")
}