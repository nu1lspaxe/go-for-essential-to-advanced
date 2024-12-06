package algorithm

type InterleavingIterator[T any] struct {
	iterators []([]T)
	indices   []int // Current indices for each iterators
}

func NewInterleavingIterator[T any](iterators []([]T)) *InterleavingIterator[T] {
	return &InterleavingIterator[T]{
		iterators: iterators,
		indices:   make([]int, len(iterators)),
	}
}

func (it *InterleavingIterator[T]) HasNext() bool {
	for i, idx := range it.indices {
		if idx < len(it.iterators[i]) {
			return true
		}
	}
	return false
}

func (it *InterleavingIterator[T]) Next() (T, bool) {
	var val T
	for i := 0; i < len(it.iterators); i++ {
		if it.indices[i] < len(it.iterators[i]) {
			val = it.iterators[i][it.indices[i]]
			it.indices[i]++
			return val, true
		}
	}
	return val, false
}
