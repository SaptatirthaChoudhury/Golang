package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mutex sync.Mutex // Mutex to protect the counter
}

func (c *Counter) Increment() {
	c.mutex.Lock()         // Acquire lock
	defer c.mutex.Unlock() // Ensure lock is released
	c.count++              // Critical section: increment counter
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	// Launch 100 goroutines to increment the counter

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 10000 {
				counter.Increment() // Safely increment shared counter
			}
		}()

	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final counter value:", counter.count)
}
