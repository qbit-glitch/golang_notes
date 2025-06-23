package main

import (
	"fmt"
	"time"
)

/*
- Goroutines are just functions that leave the main thread and run in the 
background and come back to join the main thread once the functions
are finished/ready to return any value

- Goroutines do not stop the program flow and are non-blocking
*/



func sayHello(){
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}


func main() {

	var err error

	fmt.Println("Starting the program")
	go sayHello()
	fmt.Println("After sayHello function")

	// err =  doWork()

	go func(){
		err = doWork()
	}()

	go printLetters()
	go printNumbers()
	time.Sleep(2 * time.Second)

	if err!= nil{
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully")
	}

}

func printNumbers(){
	for i:=0; i<5; i++ {
		fmt.Println("Number:",i, "Time:", time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters(){
	for _, letter := range "abcde"{
		fmt.Println("Letter:",string(letter), "Time:", time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}


func doWork() error{
	// Simulating work
	time.Sleep(1 * time.Second)

	// simulate the error
	return fmt.Errorf("error while working")
}

