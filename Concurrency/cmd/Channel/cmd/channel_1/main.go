package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	userID := 10
	resChannel := make(chan string)

	go fetchUserData(userID, resChannel)
	go fetchUserRecommendations(userID, resChannel)
	go fetchUserLikes(userID, resChannel)

	// fmt.Println(userData)
	// fmt.Println(userRecs)
	// fmt.Println(userLikes)

	for resp := range resChannel {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))
}

func fetchUserData(_ int, resChannel chan string) {
	time.Sleep(80 * time.Millisecond)

	resChannel <- "user data"
}

func fetchUserRecommendations(_ int, resChannel chan string) {
	time.Sleep(120 * time.Millisecond)

	resChannel <- "user recommendations"
}

func fetchUserLikes(_ int, resChannel chan string) {
	time.Sleep(50 * time.Millisecond)

	resChannel <- "user likes"
}
