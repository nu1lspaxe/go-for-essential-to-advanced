package goroutine

import (
	"fmt"
	"sync"
)

type Reporter struct {
	worker int
	msgs   chan string
	wg     sync.WaitGroup
	closed chan struct{}
	once   sync.Once
}

func NewReporter(worker, buffer int) *Reporter {
	return &Reporter{
		worker: worker,
		msgs:   make(chan string, buffer),
		closed: make(chan struct{}),
	}
}

func (r *Reporter) Run(stop <-chan struct{}) {
	go func() {
		<-stop
		fmt.Println("stop...")
		r.shutdown()
	}()

	for i := 0; i < r.worker; i++ {
		r.wg.Add(1)
		go func() {
			defer r.wg.Done()
			for {
				select {
				case <-r.closed:
					return
				case msg := <-r.msgs:
					fmt.Printf("report: %s\n", msg)
				}
			}
		}()
	}

	r.wg.Wait()
	fmt.Println("report workers exit...")
}

func (r *Reporter) shutdown() {
	r.once.Do(func() {
		close(r.closed)
	})
}

func (r *Reporter) Report(data string) {
	select {
	case <-r.closed:
		fmt.Printf("reporter is closed, data will be discarded: %s \n", data)
	default:
	}

	select {
	case <-r.closed:
		fmt.Printf("reporter is closed, data will be discarded: %s \n", data)
	case r.msgs <- data:
	}
}
