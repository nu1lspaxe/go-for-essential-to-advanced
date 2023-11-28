package concurrency

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func MultiChannel() {
	outChan := make(chan int)
	errChan := make(chan error)
	finishChan := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(outChan chan<- int, errorChan chan<- error, val int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			outChan <- val
			if val == 60 {
				errChan <- errors.New("error in 60")
			}
		}(outChan, errChan, i, &wg)
	}

	go func() {
		wg.Wait()
		fmt.Println("Finish all jobs.")
		close(finishChan)
	}()

LOOP:
	for {
		select {
		case val := <-outChan:
			fmt.Println("Finished: ", val)
		case err := <-errChan:
			fmt.Println("Error: ", err)
			break LOOP
		case <-finishChan:
			break LOOP
		case <-time.After(100000 * time.Millisecond):
			break LOOP
		}
	}
}
