package ctxWith

import (
	"context"
	"fmt"
)

const endNum = 10

func generate(ctx context.Context) <-chan int {
	numsCh := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case numsCh <- n:
				n++
			}
		}
	}()
	return numsCh
}

func RunWithCancel() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	for n := range generate(ctx) {
		fmt.Println(n)
		if n == endNum {
			break
		}
	}
}
