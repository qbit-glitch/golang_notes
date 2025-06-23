package main

import (
	"fmt"
	"time"
)

func main() {

	// simpleClosingChannel()
	// receiveFromClosedChannel()
	// rangeOverClosedChannel()
	// runtimePanicInAction()

	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println("Received:", val)
	}
}


func receiveFromClosedChannel(){
	ch := make(chan int)
	close(ch)

	val, ok := <- ch
	if !ok {
		fmt.Println("Channel is Closed ...")
		return
	}
	fmt.Println("Received:", val)
}

func simpleClosingChannel(){
	ch := make(chan int)

	go func(){
		for i:=range 5{
			ch <- i
		}
		close(ch)
	}()

	for value := range ch {
		fmt.Println("Received:", value)
	}
}

func rangeOverClosedChannel(){
	ch := make(chan int)
	go func(){
		for i := range 5{
			ch <- i
		}
		close(ch)
	}()
	for val := range ch {
		fmt.Println("Received Value:", val)
	}
}

func runtimePanicInAction(){
	ch := make(chan int)

	go func(){
		close(ch)
		// close(ch)
	}()
	time.Sleep(time.Second)
}


func producer(ch chan <- int){
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

func filter(in <- chan int, out chan <- int){
	for i := range in {
		if i%2 != 0 {
			out <- i
		}
	}
	close(out)
}
