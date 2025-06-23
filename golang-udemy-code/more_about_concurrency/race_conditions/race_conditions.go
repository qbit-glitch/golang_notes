package main

import (
	"fmt"
	"sync"
)

func main() {
	mutexStructMain()

}

type counter struct {
	mu sync.Mutex
	count int
}

func (c *counter) increment(){
	// c.mu.Lock()
	// defer c.mu.Unlock()
	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func mutexStructMain(){
	var wg sync.WaitGroup
	counter := &counter{}

	numGoroutines := 100

	for range numGoroutines{
		wg.Add(1)
		go func(){
			defer wg.Done()
			for range 1_000_000{
				counter.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Final Value of counter:", counter.count)
}
