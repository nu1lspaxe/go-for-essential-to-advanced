package mutex

import (
	"sync"
	"testing"
)

func MutexBenchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

// RWMutex : Mutex = 10 : 1
func BenchmarkReadMore(b *testing.B) {
	MutexBenchmark(b, &Lock{}, 9, 1)
}

func BenchmarkReadMoreRW(b *testing.B) {
	MutexBenchmark(b, &RWLock{}, 9, 1)
}

// RWMutex : Mutex = 1 : 1

func BenchmarkWriteMore(b *testing.B) {
	MutexBenchmark(b, &Lock{}, 1, 9)
}

func BenchmarkWriteMoreRW(b *testing.B) {
	MutexBenchmark(b, &RWLock{}, 1, 9)
}

// RWMutex : Mutex = 2 : 1
func BenchmarkEqual(b *testing.B) {
	MutexBenchmark(b, &Lock{}, 5, 5)
}

func BenchmarkEqualRW(b *testing.B) {
	MutexBenchmark(b, &RWLock{}, 5, 5)
}
