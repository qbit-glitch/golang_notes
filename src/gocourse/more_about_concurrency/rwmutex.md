# `RWMutex`

## Introduction

RWMutex stands for read-write mutex, is a synchronization primitive in Go that allows multiple readers to hold the lock simultaneously while ensuring exclusive access for a single writer. It provides an efficient way to handle concurrent read and write operations, particularly when read operations are frequent and writes are infrequent.

RWMutex is designed to optimize scenarios where multiple goroutines need to read shared data concurrently. But write operations are less frequent.

So RWMutex helps to improve performance by reducing contention during read operations while still maintaining exclusive access for write operations.

**Key Concepts of sync.RWMutex**
    
- Read Lock (RLock): allows multiple goroutine to acquire RLock simultaneously. It is used when a goroutine needs to read shared data without modifying it. 

- Write Lock (Lock): ensures exclusive access to the shared resources and only one goroutine can hold the write lock at a time. Moreover all readers and writers are blocked until the write block is released.

- Unlock (Unlock and RUnlock)

**When to use RWMutex**
- Read Heavy Workloads
- Shared Data Structures

## Code
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwmu sync.RWMutex
	counter int
)

func readCounter(wg *sync.WaitGroup){
	defer wg.Done()
	rwmu.RLock()
	fmt.Println("Read Counter:", counter)
	rwmu.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, value int){
	defer wg.Done()
	rwmu.Lock()
	counter = value
	fmt.Println("Writing value to counter: Done")
	rwmu.Unlock()
}


func main() {
	var wg sync.WaitGroup
	for range 5{
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	time.Sleep(3*time.Second)
	go writeCounter(&wg, 18)

	wg.Wait()

}
```



**How RWMutex Works**
- Read Lock Behavior
- Write Lock Behavior
- Lock Contention and Starvation


- When a write lock is requested, may block readers if a write lock is pending. Conversely long held read locks can delay the acquisition of a write lock. Only one goroutine can acquire the write lock at a time. 

- While a goroutine holds the write lock, no other goroutine can acquire either a read or write lock. However for the read lock behavior, multiple goroutines can acquire the read lock simultaneously, provided no go routine holds the write.

- Read Locks are shared and do not block other readers.

- Starvatinos means that your write operation (or any other operation) needs to acquire the lock but it is being held in a limbo, waiting for the lock to be released.

**Best Practices for Using RWMutex**
- Minimize Lock Duration: to avoid blocking other goroutines unnecessarily.

- Avoid Lock Starvation: Be mindful of long held read locks potentially causing write lock starvation. If write operations are critical, ensure that read operations are not indefinitely blocking writes because then your write operation will be starving.

- Avoid Deadlocks
- Balance Read and Write Operations

**Advanced Use Cases:**
- Caching with RWMutex
- Concurrent Data Structures

