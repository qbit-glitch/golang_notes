# Race Conditions

## Introduction
A race condition occurs when the outcome of a program depends on the relative timing of uncontrollable events such as thread or goroutine scheduling. It usually happens when multiple threads or goroutines access shared resources concurrently without proper synchronizatino, leading to unpredictable and incorrect behavior.

**Why does it matter ?**
- Race conditions can cause bugs that are difficult to reproduce and debug, leading to unreliable and inconsistent program behavior.

## Code:

```go
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
	// c.mu.Lock()     // --> Possible Solution is using mutexes
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

```

## Notes

- To check if the program has a race condition, add the `-race` flag during running the program.
    
```go
go run -race race_conditions.go
```

- When we have multiple goroutines accessing the same value, trying to modify the same value or trying to do something at the same time with a same type/variable/object, then in that case, use this `-race` flag and find if you have a data race in your program.

- Go provides a builtin race detector tool that helps identify the race conditions in your programs. The race detector monitors accesses to shared variables and reports data races during execution. Finally in the output, the race detector shows where data races occur including the read and write operations.

- We use mutexes or stateful goroutines or atomic operations to avoid race conditions.

- **Best Practices to Avoid Race Conditions :**
    - Proper Synchronization: use synchronization primitives like mutexes or atomic operations to ensure exclusive access to shared resources.

    - Minimize Shared State: reduce the amount of shared state between concurrent operations to lower the risk of race conditions.

    - Encapsulate State: use encapsulation to manage state within structs or functions limiting exposure to shared data.

    - Code Reviews and Testing: regularly review code for potential race conditions and utilize tools like the race detector to identify issues during development.

- **Practical Considerations**
    - Complexity of Synchronization
    - Avoiding Deadlocks
    - Performance Impact

