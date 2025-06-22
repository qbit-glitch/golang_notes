# `sync.Once`

## Intro

A once ensures that a piece of code is executed only once, regardless of how many goroutines attempt to execute it. It is useful for initializing shared resources or performing setup tasks.

## Code
```go
package main

import (
	"fmt"
	"sync"
)


var once sync.Once

func initialize(){
	fmt.Println("This function is executed only once, no matter how many times you call it")
}

func main() {
	var wg sync.WaitGroup
	for i:= range 10{
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println("Goroutine: #", i)
			once.Do(initialize)
		}()
	}
	wg.Wait()
}
```