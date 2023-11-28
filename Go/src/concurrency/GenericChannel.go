package concurrency

import (
	advanced "advanced"
	"fmt"
)

func channelGenerate[T advanced.Int32AndFloat32]() chan T {
	ch := make(chan T)
	return ch
}

func sum[T advanced.Int32AndFloat32](ch chan T, nums []T) {
	var result T
	for _, num := range nums {
		result += num
	}
	ch <- result
}

func GenericChannel() {
	numsInt := []int32{1, 2, 3, 4, 5}
	numsFloat := []float32{0.3, 1.8, 3.3, 4.6, 5.0}

	chInt := channelGenerate[int32]()
	chFloat := channelGenerate[float32]()

	go sum(chInt, numsInt)
	go sum(chFloat, numsFloat)

	fmt.Println("sumsInt: ", <-chInt)
	fmt.Println("sumFloat: ", <-chFloat)
}
