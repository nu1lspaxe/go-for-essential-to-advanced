/*
Channel with capacity is "Buffered Channel", which is asynchronous, while without capacity one called "Unbuffered Channel", which is synchronous.

In a pipeline:
start -> [stage 1] -> [stage 2] -> [stage 3] -> end
*/

package concurrency

import "fmt"

// return a read-only channel (data channel)
func intSlice2Ch(nums []int) <-chan int {
	// note: unbuffered channel means it doesn't matter that return out but go routine is still running.
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out) // close channel after send all data
	}()

	return out
}

// receive a read-only channel (data channel)
// return a read-only channel (final channel)
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func RunPipeline() {
	nums := []int{2, 3, 10, 8, 6}

	// Stage 1: dataChannel, let input into data channel
	dataCh := intSlice2Ch(nums)

	// Stage 2: finalChannel, let data in data channel do something. (Use square number here)
	finalCh := square(dataCh)

	// Stage 3: Print result in final channel
	for result := range finalCh {
		fmt.Println(result)
	}
}
