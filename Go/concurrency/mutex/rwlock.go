package mutex

import (
	"sync"
	"time"
)

// The sync.RWMutex, on the other hand, allows multiple concurrent
// read operations while locking the resource exclusively during
// write operations. This makes it suitable for scenarios where read
// operations significantly outnumber write operations.

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (l *RWLock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *RWLock) Read() {
	l.mu.RLock()
	_ = l.count
	time.Sleep(cost)
	l.mu.RUnlock()
}
