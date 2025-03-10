package main

import (
	"fmt";
)

func send(channel chan int, number int) {
	channel <- number
}

func print(channel chan int) int {
	return <-channel
}

func main() {
	channels := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
	}

	go send(channels[0], 0)

	for i := 0; i < 9; i++ {
		go send(channels[i+1],<-channels[i]+1)
	}
	
	fmt.Println(print(channels[9]))
}