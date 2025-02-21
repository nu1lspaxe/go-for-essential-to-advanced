package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Process(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second)
	ch <- int(time.Now().Unix())
}

func CreateChannels() {
	channels := make([]chan int, 10)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		wg.Add(1)
		go Process(channels[i], &wg)
	}

	for idx, ch := range channels {
		<-ch
		fmt.Printf("Routine: %v quit.\n", idx)
	}

	wg.Wait()
}
