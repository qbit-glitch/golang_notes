package main

import (
	"fmt"
	"time"
)

type StatefulWorker struct{
	count int
	ch chan int
}

func (w *StatefulWorker) Start(){
	go func(){
		for{
			select{
			case value := <- w.ch:
				w.count += value
				fmt.Println("Current Count:", w.count)
			}
		}
	}()
}

func (w *StatefulWorker) Send(value int){
	w.ch <- value
}

func main() {

	stworker := StatefulWorker{
		ch: make(chan int),
	}
	stworker.Start()

	for i := range 10 {
		stworker.Send(i)
		time.Sleep(500 * time.Millisecond)
	}

}