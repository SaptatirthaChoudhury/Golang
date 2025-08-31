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

func printSum(v1, v2 *value) {
	defer wg.Done()

	// Order locks by memory address (pointer comparison)
	if v1 == v2 {
		v1.mu.Lock()
		defer v1.mu.Unlock()
	} else if v1.value < v2.value {
		v1.mu.Lock()
		v2.mu.Lock()
		defer v2.mu.Unlock()
		defer v1.mu.Unlock()
	} else {
		v2.mu.Lock()
		v1.mu.Lock()
		defer v1.mu.Unlock()
		defer v2.mu.Unlock()
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("sum=%v\n", v1.value+v2.value)
}

func main() {
	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
