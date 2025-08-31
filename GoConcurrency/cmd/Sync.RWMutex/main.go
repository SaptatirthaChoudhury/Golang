package main

import (
	"fmt"
	"sync"
	"time"
)

var rw sync.RWMutex
var count int

func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rw.RLock()
	fmt.Printf("Reader %d: count=%d\n", id, count)
	time.Sleep(100 * time.Millisecond)
	rw.RUnlock()
}

func writer(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rw.Lock()
	count++
	fmt.Printf("Writer %d updated count=%d\n", id, count)
	time.Sleep(200 * time.Millisecond)
	rw.Unlock()
}

func main() {
	var wg sync.WaitGroup

	// spawn readers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go reader(i, &wg)
	}

	// spawn writers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go writer(i, &wg)
	}

	wg.Wait()
}
