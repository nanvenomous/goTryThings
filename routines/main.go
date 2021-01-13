package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	messages := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go one(messages)
	wg.Add(1)
	go two(messages)
	wg.Wait()
}

func one(messages chan<- string) {
	for {
		time.Sleep(time.Second)
		messages <- "hey"

	}
}

func two(messages chan string) {
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
