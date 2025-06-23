package main

import (
	"fmt"
	"time"
)

func multiplexingMultipleChannels() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1

	}()

	go func() {

		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()

	// We have to range over the same number channels 
	for range 2 {
		select {
		case msg := <-ch1:
			fmt.Println("Received from ch1:", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch2:", msg)
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("No Channels ready...")
		}
	}

	fmt.Println("End of Program")

}


func timeouts(){
	ch := make(chan int)

	go func(){
		time.Sleep(2 * time.Second)
		ch <- 2
		close(ch)
	}()

	select {
	case msg:= <- ch :
		fmt.Println("Received:", msg)
	case <- time.After(3 * time.Second):
		fmt.Println("Timeout!")
	}
}


func receiveMessages(){
	ch := make(chan int)

	go func(){
		ch <- 1
		close(ch)
	}()

	for {
		select {
		case msg, ok := <- ch :
			// ok will return false only when the channel is closed and empty.
			if !ok {
				fmt.Println("Channel Closed!")
				// Clean up activities
				return
			}
			fmt.Println("Received:", msg)
		}
	}
}


func main(){
	timeouts()
	receiveMessages()
}