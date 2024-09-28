package concurrency

import (
	"sync/atomic"
	"unsafe"
)

// non-lock stack
type LFStack struct {
	head unsafe.Pointer
}

type Node struct {
	val  int32
	next unsafe.Pointer
}

func NewLFStack() *LFStack {
	n := unsafe.Pointer(&Node{})
	return &LFStack{head: n}
}

func (s *LFStack) Push(v int32) {
	n := &Node{val: v}

	for {
		old := atomic.LoadPointer(&s.head)
		n.next = old
		if atomic.CompareAndSwapPointer(&s.head, old, unsafe.Pointer(n)) {
			return
		}
	}
}

func (s *LFStack) Pop() int32 {
	for {
		old := atomic.LoadPointer(&s.head)
		if old == nil {
			return 0
		}

		oldNode := (*Node)(old)
		next := atomic.LoadPointer(&oldNode.next)
		if atomic.CompareAndSwapPointer(&s.head, old, next) {
			return oldNode.val
		}
	}
}
