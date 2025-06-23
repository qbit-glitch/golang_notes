package main

import (
	"fmt"
	"sync"
	"time"
)


type LeakyBucket struct{
	capacity int
	leakRate time.Duration
	tokens int
	lastLeak time.Time
	mu sync.Mutex
}

func NewRateLimiter(capacity int, leakRate time.Duration) *LeakyBucket{
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens: capacity,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)
	tokensToAdd := int(elapsedTime / lb.leakRate)
	lb.tokens += tokensToAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)

	/*
		lb.lastLeak = lb.lastLeak.Add(elapsedTime)    // WRONG APPROACH
		elapsedTime = 1.3 seconds
		initial lastLeak = 0
		tokensToAdd = int(1.3 / 0.5) = int(2.6) = 2
		lb.lastLeak = lb.lastLeak + time.Duration(tokensToAdd) * lb.leakRate = 0 + (2 * 0.5) = 1 sec
		lb.lastLeak = lb.lastLeak + elapsedTime = 0 + 1.3   // wasting the resources for 0.3 seconds extra
	*/

	fmt.Printf("Tokens added %d | Tokens subtracted %d | Total Tokens: %d\n", tokensToAdd, 1, lb.tokens)
	fmt.Println("Last leak time: ", lb.lastLeak)

	if lb.tokens > 0{
		lb.tokens--
		return true
	}

	return false
}

func main() {

	leakyBucketInst := NewRateLimiter(5, 500 * time.Millisecond)
	var wg sync.WaitGroup

	for range 10{
		wg.Add(1)
		go func(){
			defer wg.Done()
			if leakyBucketInst.Allow(){
				fmt.Println("Current Time:", time.Now())
				fmt.Println("Request Allowed")
			} else {
				fmt.Println("Current Time:", time.Now())
				fmt.Println("Request Denied --- XXX")
			}
			time.Sleep(200 * time.Millisecond)
		}()
	}
	wg.Wait()
	fmt.Println("Program execution finished...")

}