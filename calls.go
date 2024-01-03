package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	now := time.Now()
	ws := &sync.WaitGroup{}
	playerId := 10
	playerch := make(chan string, 256)
	// goroutines
	go fetchPlayerStats(playerId, playerch, ws)
	go fetchPlayerHistory(playerId, playerch, ws)
	go fetchPlayerFriends(playerId, playerch, ws)
	//add the above 2 rotuinnes to waitgroup
	ws.Add(3)
	//wait for the sloweset fetch
	ws.Wait()
	//close channel else deadlock
	close(playerch)

	for data := range playerch {
		fmt.Println(data)
	}

	fmt.Println(time.Since(now))
}

func fetchPlayerStats(id int, playerch chan string, ws *sync.WaitGroup) {
	time.Sleep(60 * time.Millisecond)
	playerch <- "shriahri sudevan"
	ws.Done()
}

func fetchPlayerHistory(id int, playerch chan string, ws *sync.WaitGroup) {
	//this is the slowest fetch so in an async env total time == slowest func
	// if it was a sync env then total time == sum of all exec times
	time.Sleep(180 * time.Millisecond)
	playerch <- "lost"
	ws.Done()
}

func fetchPlayerFriends(id int, playerch chan string, ws *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)
	playerch <- "Shreyas"
	ws.Done()
}
