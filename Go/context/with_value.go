package context

import (
	"context"
	"fmt"
	"time"
)

type TraceCode string

func ValueWorker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code.")
	}

LOOP:
	for {
		fmt.Printf("worker, trace code: %s\n", traceCode)
		time.Sleep(time.Millisecond * 10)

		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func RunWithValue() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*50)
	defer cancel()

	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "000000")
	wg.Add(1)
	go ValueWorker(ctx)
	wg.Wait()
	fmt.Println("work over.")
}
