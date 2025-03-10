package main

import (
	"fmt";
	"math/rand";
	"time";
	"sync";
)

var wg sync.WaitGroup

func chain(channel chan int, out_channel chan int, T int, id int) {
	defer wg.Done()
	number := 0

	for i := 0; i < T; i++ {
		number = <-channel
		fmt.Println("Ouvrier", id, "travaille sur", number)

		if(id == 4 && i == 4){
			wg.Done()
		}

		out_channel <- number+1
	}
}

func main(){
	rand.Seed(time.Now().UnixNano())

	N := 5

	T := 5

	channels := make([]chan int, N)

	for i := 0; i < N; i++ {
		channels[i] = make(chan int)
	}

	for i := 0; i < N; i++ {
		wg.Add(1)

		if(i == N-1){
			go chain(channels[i], channels[0], T, i)
		}else{
			go chain(channels[i], channels[i+1], T, i)
		}
	}

	channels[0] <- 0

	wg.Wait()

	fmt.Println(<-channels[0])
}