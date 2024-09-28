package threadpools

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

const (
	jobsCount    = 10
	workersCount = 2
)

func TestWorkerPool(t *testing.T) {
	wp := New(workersCount)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	go wp.GenerateFrom(runJobs())
	go wp.Run(ctx)

	for {
		select {
		case result, ok := <-wp.Results():
			if !ok {
				continue
			}

			i, err := strconv.ParseInt(string(result.Descriptor.ID), 10, 64)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			val := result.Value.(int)
			if val != int(i)*2 {
				t.Fatalf("wrong value: %v; expected: %v", val, int(i)*2)
			}
		case <-wp.Done:
			return
		}
	}
}

func TestWorkerPool_TimeOut(t *testing.T) {
	wp := New(workersCount)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Nanosecond*10)
	defer cancel()

	go wp.Run(ctx)

	for {
		select {
		case result := <-wp.Results():
			if result.Err != nil && result.Err != context.DeadlineExceeded {
				t.Fatalf("expected error: %v; got: %v", context.DeadlineExceeded, result.Err)
			}
		case <-wp.Done:
			return
		}
	}
}

func TestWorkerPool_Cancel(t *testing.T) {
	wp := New(workersCount)

	ctx, cancel := context.WithCancel(context.TODO())
	cancel() // 測試 cancel() 所以不需要 defer

	go wp.Run(ctx)

	for {
		select {
		case result := <-wp.Results():
			if result.Err != nil && result.Err != context.Canceled {
				t.Fatalf("expected error: %v; got: %v", context.Canceled, result.Err)
			}
		case <-wp.Done:
			return
		}
	}
}

func runJobs() []Job {
	jobs := make([]Job, jobsCount)

	for i := 0; i < jobsCount; i++ {
		jobs[i] = Job{
			Descriptor: JobDescriptor{
				ID:       JobID(fmt.Sprintf("%v", i)),
				JType:    "anyType",
				MetaData: nil,
			},
			ExecFn: execFn,
			Args:   i,
		}
	}
	return jobs
}
