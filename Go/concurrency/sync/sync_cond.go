package sync

import (
	"log"
	"sync"
	"time"
)

// Each Cond has an associated Locker L (*Mutex or *RWMutex),
// which must be held when changing the condition and
// when calling the Wait method.
//
// func NewCond(l Locker) *Cond

// Broadcast wakes all goroutines waiting on c.
// func (c *Cond) Broadcast()

// Signal wakes one goroutine waiting on c
// func (c *Cond) Signal()

// Wait atomically unlocks c.L and suspends execution
// of the calling goroutine. After later resuming
// execution,Wait locks c.L before returning.
//
// 		c.Lock()
// 		for !condition() {
// 			c.Wait()
// 		}
// 		// make use of condition
// 		c.L.Unlock()
//
// func (c *Cond) Wait()

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func RunCond() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
}
