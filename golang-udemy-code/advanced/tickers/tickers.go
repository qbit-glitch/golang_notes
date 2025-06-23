package main

import (
	"fmt"
	"time"
)

func main() {
	// tickerBasics()
	// periodicTaskExecution()
	tickerAndTimer()

}

func tickerBasics(){
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	// for tick := range ticker.C {
	// 	fmt.Println("Tick at:", tick)
	// }
	var i int = 1
	for range ticker.C {
		i *= 2
		fmt.Println(i)
	}
}



// --------- SCHEDULING LOGGING PERIODIC TASKS, POLLING FOR UPDATES -------
func periodicTask(){
	fmt.Println("Performing Periodic task at:", time.Now())
}

func periodicTaskExecution(){
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select{
		case <-ticker.C:
			periodicTask()
		}
	}
}


// 
func tickerAndTimer(){
	ticker := time.NewTicker(500 * time.Millisecond)
	stop := time.After(5 * time.Second)

	for{
		select{
		case tick := <-ticker.C :
			fmt.Println("Tick at:", tick)
		case <- stop :
			fmt.Println("Stopping Ticker..")
			return
		}
	}
}