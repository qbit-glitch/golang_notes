package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	atomicMain()

}

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increment(){
	atomic.AddInt64(&ac.count, 2)
}

func (ac *AtomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func atomicMain(){
	var wg sync.WaitGroup
	numGoroutines := 10
	counter := &AtomicCounter{}
	value := 0
	for range numGoroutines{
		wg.Add(1)
		go func(){
			defer wg.Done()
			for range 1000 {
				counter.increment()
				value++
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", counter.count)
	fmt.Printf("Final Value: %d\n", value)
}