# sync.NewCond

## Introduction

NewCond is a function in Go's sync package that creates a new condition variable. A condition variable is a synchronization primitive that allows goroutines to wait for certain conditions to be met while holding a lock. It is used to signal one ore more goroutines that some condition has changes.

Condition variables are essential for more complex synchronization scenarios beyond simple locking mechanisms. They are useful in situations where goroutines need to wait for specific conditions or events before proceeding.

- Key Concepts of `sync.NewCond`:
    - Condition Variables
    - Mutex and Condition Variables

- Methods of `sync.Cond`
    - `Wait()`
    - `Signal()`
    - `Broadcast()`


## Code

```go
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
```

## Notes:

### Key Points :

- Signal is for waking up the other goroutine. Wait is for making our goroutine fall asleep.

- `sync.NewCond`: it allows goroutines to wait for or signal changes in program state. It creates a new condition variable associate with the buffers mutex, which it takes as an argument.

- `b.cond.Wait()`: makes the goroutine wait until the signal is received. It puts the goroutines to sleep and Signal wakes up that sleeping goroutine.

- `b.cond.Signal()`: sends a notification to notify a consumer.

### Best Practices for using `sync.NewCond`

- Ensure Mutex is held
- Avoid spurious wakeups
- Use condition variables judiciously
- Balance signal and broadcast

### Advanced Use Cases
- Task Scheduling
- Resource Pools
- Event Notification Systems
