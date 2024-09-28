package concurrency

import (
	"fmt"
	"time"
)

func Process(ch chan int) {
	time.Sleep(time.Second)
	ch <- int(time.Now().Unix())
}

func CreateChannels() {
	channels := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go Process(channels[i])
	}

	for idx, ch := range channels {
		<-ch
		fmt.Printf("Routine: %v quit.\n", idx)
	}
}
