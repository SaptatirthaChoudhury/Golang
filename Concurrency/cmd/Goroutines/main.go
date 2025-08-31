package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	time.Sleep(time.Second * 2)

	fmt.Println("hello!")

	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "data_2"
	}()

	//msg := <-myChannel
	//fmt.Println(msg)

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
    // Buffered channel with capacity 3
	charChannel := make(chan string, 3)
	char := []string{"a", "b", "c"}

	for _, s := range char {
		charChannel <- s
	}

	close(charChannel) // closing channel

	for result := range charChannel {
		fmt.Println(result)
	}


}
