package main

import (
	"fmt"
	"time"
)

func main() {
	c := make
	go count("sheep", c)

	msg := <-c // To recieve msg from channel. Sending and Receiving are blocking operations.
	fmt.Println(msg)
}

func countUsingChannels(thing string, c chan string) {
	for i := 0; i <= 5; i++ {
		c <- thing // Send 'thing' over the channel.
		time.Sleep(time.Millisecond * 500)
	}
}
