package main

import (
	"fmt"
	"time"
)

func main() {
	nonBlockingReceiveOperation()
	nonBlockingSendOperation()
	nonBlockingOperationInRealTimeSystem()
}


func nonBlockingReceiveOperation(){
	// ------ NON BLOCKING RECEIVE OPERATION ------
	ch := make(chan int)

	select {
	case msg := <- ch :
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No messages available.")
	}
}

func nonBlockingSendOperation(){
	// ----- NON BLOCKING SEND OPERATION -----
	ch := make(chan int)

	select{
	case ch <- 2:
		fmt.Println("Sent message.")
	default:
		fmt.Println("Channel is not ready to receive.")
	}
}

func nonBlockingOperationInRealTimeSystem(){
	// ----- NON BLOCKING OPERATION IN REALTIME SYSTEM ----
	data := make(chan int)
	quit := make(chan bool)

	// Consumer : checking for data after every 0.5s
	go func(){
		for {
			select {
			case d:= <- data:
				fmt.Println("Data received:", d)
			case <- quit :
				fmt.Println("Stopping ...")
				return

			default: 
				fmt.Println("Waiting for data ..")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()


	// Producer
	for i := range 10 {
		data <- i
		time.Sleep(1 * time.Second)
	}
	quit <- true
}