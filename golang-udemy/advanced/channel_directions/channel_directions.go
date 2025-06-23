package main

import "fmt"

func main() {

	simpleSendAndReceive() 

	ch := make(chan int)
	producer(ch)
	consumer(ch)
}

func simpleSendAndReceive(){
	ch := make(chan int)

	go func(ch chan <- int){
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}(ch)

	for value := range ch {
		fmt.Println("Received:", value)
	}
}


// Receive Only Channel
func consumer(ch <- chan int){

	for value := range ch {
		fmt.Println("Received:", value)
	}
}


// SEND ONLY Channel
func producer(ch chan <- int){
	go func(ch chan <- int){
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}(ch)
}


/* Producer Consumer scenario

This is a producer where we create a producer that can only
produce data. And hence it accepts a send-only channel.

And then a consumer, which receives data and hence it accepts
a receive-only channel.
*/