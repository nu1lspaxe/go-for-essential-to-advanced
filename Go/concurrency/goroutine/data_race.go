package goroutine

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// Goroutine with temp var
func DataRaceCase1() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		// Use variable as args
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// Shared var in parallel
func DataRaceCase2(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("/tmp/file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// cannot use _, err = f1.Write(data)
			// it could cause data race due to
			// share error with main goroutine
			_, err := f1.Write(data)
			res <- err
			f1.Close()
		}()
	}

	f2, err := os.Create("/tmp/file2")
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err := f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}

// Data race case 3
// sync.Mutex to manager error

var (
	service map[string]string
	servMux sync.Mutex
)

func RegisterServ(name, addr string) {
	servMux.Lock()
	defer servMux.Unlock()
	service[name] = addr
}

func LookupServ(name string) string {
	servMux.Lock()
	defer servMux.Unlock()
	return service[name]
}

// Data race case 4
type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	// Conflict
	// w.last = time.Now().UnixNano()
	atomic.StoreInt64(&w.last, time.Now().UnixNano())
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)

			// if w.last < time.Now().Add(-10*time.Second).UnixNano()
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seoncs. Dying.")
				os.Exit(1)
			}
		}
	}()
}
