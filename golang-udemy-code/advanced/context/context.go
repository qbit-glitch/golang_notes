package main

import (
	"context"
	"fmt"
	"log"

	"time"
)

func main() {

	// differenceTODOvsBackground()
	// checkEvenOddMain()
	doWorkMain() 

}

func differenceTODOvsBackground() {
	todoContext := context.TODO()

	ctx := context.WithValue(todoContext, "name", "John")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	contextBkg := context.Background()
	ctx1 := context.WithValue(contextBkg, "city", "New York")
	fmt.Println(ctx1)
	fmt.Println(ctx1.Value("city"))

}

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Operation Cancelled"
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even", num)
		} else {
			return fmt.Sprintf("%d is odd", num)
		}
	}
}

func checkEvenOddMain() {
	ctx := context.TODO()

	result := checkEvenOdd(ctx, 5)
	fmt.Println("Result with context.TODO():", result)

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	result = checkEvenOdd(ctx, 10)
	fmt.Println("Result with context.WithTimeout():", result)

	time.Sleep(3 * time.Second)
	result = checkEvenOdd(ctx, 15)
	fmt.Println("Result with context.WithTimeout():", result)
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work Cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(800 * time.Millisecond)
	}
}

func doWorkMain() {
	rootctx := context.Background()

	// ctx, cancel := context.WithTimeout(rootctx, 3*time.Second) 
	// // timer of the context starts here. You have this specfied time 
	// // duration to use this contexts. After this time duration the 
	// // context will send a cancellation signal.
	// defer cancel()

	ctx, cancel := context.WithCancel(rootctx)

	go func(){
		time.Sleep(3 * time.Second)  // simulating a heavy task. Consider this a heavy time consuming operation.
		cancel()					 // manually canceling only after the task is finished.
	}()

	ctx = context.WithValue(ctx, "requestID", "ACEG6417")

	go doWork(ctx)

	time.Sleep(4 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No request ID found.")
	}

	logWithContext(ctx, "This is a test log message")
}


func logWithContext(ctx context.Context, message string){
	requestIDVal := ctx.Value("requestID")
	log.Printf("Request ID: %v - %v", requestIDVal, message)
}
