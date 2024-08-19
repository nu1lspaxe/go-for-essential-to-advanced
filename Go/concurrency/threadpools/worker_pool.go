package threadpools

import (
	"context"
	"fmt"
	"sync"
)

func worker(ctx context.Context, wg *sync.WaitGroup, job <-chan Job, results chan<- Result) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-job:
			if !ok {
				return
			}
			results <- job.execute(ctx)
		case <-ctx.Done():
			fmt.Printf("cancelled worker. Error detail: %v\n", ctx.Err())
			results <- Result{
				Err: ctx.Err(),
			}
			return
		}
	}
}

type WorkerPool struct {
	workersCount int
	jobQueue     chan Job
	results      chan Result
	Done         chan struct{}
}

func New(wcount int) *WorkerPool {
	return &WorkerPool{
		workersCount: wcount,
		jobQueue:     make(chan Job, wcount),
		results:      make(chan Result, wcount),
		Done:         make(chan struct{}),
	}
}

func (wp *WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < wp.workersCount; i++ {
		wg.Add(1)
		go worker(ctx, &wg, wp.jobQueue, wp.results)
	}

	wg.Wait()
	close(wp.Done)
	close(wp.results)
}

func (wp *WorkerPool) Results() <-chan Result {
	return wp.results
}

func (wp *WorkerPool) GenerateFrom(jobBulk []Job) {
	for _, job := range jobBulk {
		wp.jobQueue <- job
	}
	close(wp.jobQueue)
}
