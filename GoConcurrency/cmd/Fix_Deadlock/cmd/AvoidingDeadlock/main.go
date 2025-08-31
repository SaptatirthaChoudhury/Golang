package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func printSum(v1, v2 *value) {
	defer wg.Done()

	// Ensure consistent lock order based on memory address
	first, second := v1, v2
	if uintptr(unsafe.Pointer(v1)) > uintptr(unsafe.Pointer(v2)) {
		first, second = v2, v1
	}

	first.mu.Lock()
	defer first.mu.Unlock()

	time.Sleep(2 * time.Second)
	second.mu.Lock()
	defer second.mu.Unlock()

	fmt.Printf("sum=%v\n", v1.value+v2.value)
}

func main() {
	var a, b value
	wg.Add(2)
	go printSum(&a, &b) // Goroutine 1
	go printSum(&b, &a) // Goroutine 2
	wg.Wait()
}
