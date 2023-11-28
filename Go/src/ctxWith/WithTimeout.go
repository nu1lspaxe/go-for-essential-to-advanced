package ctxWith

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func TimeoutWorker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("DB connecting...")
		time.Sleep(time.Millisecond * 10)

		select {
		case <-ctx.Done(): // ctx.WithTimeout = 50*millisecond
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func RunWithTimeout() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*50)
	defer cancel()

	wg.Add(1)
	go TimeoutWorker(ctx)
	wg.Wait()
	fmt.Println("work over.")
}
