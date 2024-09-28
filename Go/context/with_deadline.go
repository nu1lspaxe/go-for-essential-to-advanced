package context

import (
	"context"
	"fmt"
	"time"
)

func RunWithDeadline() {
	deadlline := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.TODO(), deadlline)

	defer cancel() // 若不調用，可能會使其存活時間超過必要時間

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept...")
	case <-ctx.Done(): // 其實只有這個會執行, 因為deadline=50*millisecond
		fmt.Println(ctx.Err())
	}
}
