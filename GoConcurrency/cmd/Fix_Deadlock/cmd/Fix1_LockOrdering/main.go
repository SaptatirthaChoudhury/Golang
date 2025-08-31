package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

// Always lock in a fixed order (by pointer address or ID)
// This ensures both goroutines lock a and b in the same order, no matter how you pass them.

func printSum(v1, v2 *value) {
	defer wg.Done()

	// Locking order rule: lock the lower address first
	if fmt.Sprintf("%p", v1) < fmt.Sprintf("%p", v2) {
		v1.mu.Lock()
		time.Sleep(2 * time.Second)
		v2.mu.Lock()
	} else {
		v2.mu.Lock()
		time.Sleep(2 * time.Second)
		v1.mu.Lock()
	}

	fmt.Printf("sum=%v\n", v1.value+v2.value)

	v1.mu.Unlock()
	v2.mu.Unlock()
}

func main() {
	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
