package routines

import (
	"fmt"
	"time"
)

func One(messages chan<- string) {
	for {
		time.Sleep(time.Second)
		messages <- "hey"

	}
}

func Two(messages chan string) {
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
