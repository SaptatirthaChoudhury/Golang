package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
func dowork(d time.Duration, wg *sync.WaitGroup) {
	fmt.Println("doing work ...")
	time.Sleep(d)
	fmt.Println("work is done")
	wg.Done()
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go dowork(time.Second*2, &wg)
	go dowork(time.Second*4, &wg)
	wg.Wait()
	fmt.Printf("work took %v seconds\n", time.Since(start))
}

*/

func dowork(d time.Duration, resch chan string) {
	fmt.Println("doing work ...")
	time.Sleep(d)
	fmt.Println("work is done")
	resch <- fmt.Sprintf("the result of the work %d", rand.Intn(100))
	wg.Done()
}

var wg *sync.WaitGroup

func main() {
	start := time.Now()
	resultChannel := make(chan string)
	wg = &sync.WaitGroup{}
	wg.Add(3)
	go dowork(time.Second*2, resultChannel)
	go dowork(time.Second*4, resultChannel)
	go dowork(time.Second*6, resultChannel)

	go func() {
		for res := range resultChannel {
			fmt.Println(res)
		}
		fmt.Printf("work took %v seconds\n", time.Since(start))
	}()

	wg.Wait()
	close(resultChannel)

	time.Sleep(time.Second)
}
