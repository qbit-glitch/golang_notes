package main

import (
	"fmt"
	"time"
)

func main() {
	// timerBasics()
	// longRunningOperationMain()
	// schedulingDelayedOperations()
	multipleTimers()
}


// -------- BASIC TIMER USE ------------
func timerBasics(){
	fmt.Println("Starting app...")
	
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting for timer.C")

	stopped:= timer.Stop()

	if stopped {
		fmt.Println("Timer Stopped")
	}
	timer.Reset(time.Second)
	fmt.Println("Timer reset")

	<- timer.C   // blocking in Nature
	fmt.Println("Timer expired")
}



// ----- TIMEOUT -----------------------
func longRunningOperation(){
	for i:= range 20{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func longRunningOperationMain(){
	timeout := time.After(2 * time.Second)
	done := make(chan bool)

	go func(){
		longRunningOperation()
		done <- true
	}()

	select{
	case <- timeout:
		fmt.Println("Operation Timed Out")
	case <-done :
		fmt.Println("Operation Completed")
	}
}


// ------- SCHEDULING DELAYED OPERATIONS ------------
func schedulingDelayedOperations(){
	timer := time.NewTimer(2 * time.Second)   // non-blocking timer starts

	go func(){
		<- timer.C   // receives the current time after the timer expires  
		fmt.Println("Delayed Operation Executed")
	}()

	fmt.Println("Waiting...")
	time.Sleep(3 * time.Second)     // blocking timer starts
	fmt.Println("End of Program")
	
}

 

// ---------- USING MULTIPLE TIMERS ----------------
func multipleTimers(){
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	select {
	case <- timer1.C :
		fmt.Println("Timer1 expired")
	case <- timer2.C:
		fmt.Println("Timer2 expired")
	}
}
