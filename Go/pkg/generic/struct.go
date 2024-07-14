/*
	泛型結構的方法本身不能加入泛型參數，但可以沿用結構的泛型參數，因此結構初始化時，也必須填入泛型引數。
*/

package generic

import (
	"fmt"
)

type Numbers[T Int32AndFloat32] struct {
	numbers []T
}

func (n *Numbers[T]) addNums(nums ...T) {
	n.numbers = append(n.numbers, nums...)
}

func (n *Numbers[T]) sum() T {
	var result T
	for _, num := range n.numbers {
		result += num
	}
	return result
}

func GenericStruct() {
	numsInt := []int32{1, 2, 3, 4, 5}
	numsFloat := []float32{0.3, 1.8, 3.3, 4.6, 5.0}

	numbersInt := Numbers[int32]{}
	numbersFloat := Numbers[float32]{}

	numbersInt.addNums(numsInt...)
	numbersFloat.addNums(numsFloat...)

	sumInt := numbersInt.sum()
	sumFloat := numbersFloat.sum()

	fmt.Printf("sumInt: %v (type: %T)\n", sumInt, sumInt)
	fmt.Printf("sumFloat: %v (type: %T)\n", sumFloat, sumFloat)
}
