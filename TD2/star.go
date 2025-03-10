package main

import (
	"fmt";
	"math/rand";
	"time";
	"sync";
)

var wg sync.WaitGroup

func ouvrier(channel chan int, out_channel chan int) {
	defer wg.Done()


	number := <-channel
	fmt.Println("Ouvrier travaille sur", number)

	wg.Done()

	out_channel <- number*number
}

func main(){
	rand.Seed(time.Now().UnixNano())

	N := 5

	channel := make(chan int)

	out_channels := make([]chan int, N)

	for i := 0; i < N; i++ {
		wg.Add(1)
		out_channel := make(chan int)
		out_channels[i] = out_channel
		go ouvrier(channel, out_channel)
	}

	for i := 0; i < N; i++ {
		channel <- rand.Intn(10)
	}

	wg.Wait()

	for i := 0; i < N; i++ {
		fmt.Println("Ouvrier", i, "a fini, le resultat est", <-out_channels[i])
	}
	
}