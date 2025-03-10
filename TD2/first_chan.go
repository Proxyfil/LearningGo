package main

import (
	"fmt";
)

func send(channel chan string, message string) {
	channel <- message
}

func main() {
	channel := make(chan string)

	go send(channel, "Bonjour, Goroutine !")
	
	fmt.Println(<-channel)
}