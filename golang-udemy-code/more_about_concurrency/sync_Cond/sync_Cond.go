package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

type buffer struct {
	items []int
	mu sync.Mutex
	cond *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) produce(item int){
	b.mu.Lock()
	defer b.mu.Unlock()

	// Conditional infinite for loop
	for len(b.items) == bufferSize {
		b.cond.Wait()
	}

	b.items = append(b.items, item)
	fmt.Println("Produced:", item)
	b.cond.Signal()    // signal the consumer that the producer has done it's job to produce an item.
}

func (b *buffer) consume() int{
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == 0 {
		b.cond.Wait()
		// This functions stops doing anything and waits for
		// other functions to append to the slice
	}
	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed:", item)
	b.cond.Signal()
	return item
}

func producer(b *buffer, wg *sync.WaitGroup){
	defer wg.Done()
	for i := range 10 {
		b.produce(i+1000)
		time.Sleep(200 * time.Millisecond)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup){
	defer wg.Done()
	for range 10 {
		b.consume()
		time.Sleep(1500 * time.Millisecond)
	}
}



func main() {

	buffer := newBuffer(bufferSize)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()

}