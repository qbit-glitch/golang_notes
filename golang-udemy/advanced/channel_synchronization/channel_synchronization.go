package main

import (
	"fmt"
	"time"
)

func main() {

	channelSyncFlagDemo1()
	channelSyncFlagDemo2()
	channelSyncMultipleGoRoutines()
	channelSyncDataExchange()

}

func channelSyncFlagDemo1() {
	done := make(chan struct{})

	go func() {
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	<-done
	fmt.Println("Finished ...")
}

func channelSyncFlagDemo2() {
	ch := make(chan int)

	go func() {
		ch <- 9
		time.Sleep(1 * time.Second)
		fmt.Println("Sent Value")
	}()

	value := <-ch
	fmt.Println("Value:", value)

	/*
		`Sent Value` will not be printed in the main function,
		because as soon as the value is received by the channel
		in the main function, the execution flow is so fast that the
		program does not leave a time margin for the go routine to
		execute printing the `Sent Value` statement.

		To get this printed the main thread needs to be busy doing
		something while this statement in goroutine executes. You could
		try putting the main thread to sleep for 1 second before it ends,
		to simulate some work and see it printed.
	*/
}


func channelSyncMultipleGoRoutines(){
	// SYNCHRONIZING MULTIPLE GOROUTINES AND ENSURING THAT ALL GOROUTINES ARE COMPLETE
	numGoroutines := 3
	done := make(chan int, 3)
	for i := range numGoroutines {
		time.Sleep(time.Second)
		go func(id int){
			fmt.Printf("Goroutine %d working ...\n", id)
			// time.Sleep(time.Second)
			done <- id   // SENDING SIGNAL OF COMPLETION
		}(i)
	}

	for range 2 {
		<- done      // wait for each goroutine to finish. WAIT FOR ALL GOROUTINES TO SIGNAL COMPLETION
	}

	fmt.Println("All goroutines are finished.")
}

func channelSyncDataExchange(){
	// ---- SYNCHRONIZING DATA EXCHANGE ------
	data := make(chan string)
	go func(){
		for i := range 5{
			data <- "Hello " + string(rune('0'+i))
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()
	// close(data)   // channel closed before Goroutine could send a value to the channel

	for value := range data {
		fmt.Println("Received Data:", value, ":", time.Now())
	}  // Loops over only on active channel, creates receiver each time and stops creating receiver (looping) once the channel is closed.
}
