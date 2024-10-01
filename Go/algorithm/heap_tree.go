package algorithm

import "fmt"

// MinHeap represents a Min-Heap data structure
type MinHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

// ExtractMin removes and returns the minimum element
// (the root) from the heap
func (h *MinHeap) ExtractMin() int {
	if len(h.array) == 0 {
		fmt.Println("Heap is empty")
		return -1
	}

	root := h.array[0]

	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	h.heapifyDown(0)

	return root
}

// heapifyUp restores the heap property after insertion
func (h *MinHeap) heapifyUp(index int) {
	for h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// heapifyDown restores the heap property after extraction
func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	left, right := leftChild(index), rightChild(index)
	childToCompare := 0

	// Loop while index has at least one child
	for left <= lastIndex {
		if right <= lastIndex && h.array[right] < h.array[left] {
			childToCompare = right
		} else {
			childToCompare = left
		}

		if h.array[index] <= h.array[childToCompare] {
			break
		}

		h.swap(index, childToCompare)
		index = childToCompare
		left, right = leftChild(index), rightChild(index)
	}
}

// parent returns the parent index of a given node
func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

// swap exchanges the elements at the provided indices
func (h *MinHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func MinHeapTree() {
	h := new(MinHeap)

	h.Insert(3)
	h.Insert(1)
	h.Insert(6)
	h.Insert(5)
	h.Insert(2)
	h.Insert(4)

	fmt.Println("Heap array:", h.array)

	fmt.Println("Extracted Min:", h.ExtractMin())
	fmt.Println("Extracted Min:", h.ExtractMin())
	fmt.Println("Extracted Min:", h.ExtractMin())
	fmt.Println("Heap array after extractions:", h.array)
}
