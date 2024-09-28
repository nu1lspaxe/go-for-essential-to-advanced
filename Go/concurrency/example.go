package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Task 1.
// The server can serves 10 requests per time, but now
// there are 100 requests coming

type Task interface {
	Run()
}

type TaskProcessor interface {
	Do(id int, msg string)
}

// Semaphore
type M1 struct {
	Processor TaskProcessor
}

func (m M1) Run() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 10) // limited semaphore, allow maximum 10 goroutine

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{}

			m.Processor.Do(id, fmt.Sprintf("message %d", id))

			<-sem
		}(i)
	}
	wg.Wait()
}

type DefaultProcessor struct{}

func (p DefaultProcessor) Do(id int, msg string) {
	fmt.Printf("processing message %d: %s\n", id, msg)
	time.Sleep(1 * time.Nanosecond)
}
