package advanced

import (
	"fmt"
)

func getType(item interface{}) {
	switch item.(type) {
	case int:
		fmt.Println("The type is int.")
	case string:
		fmt.Println("The type is string.")
	case bool:
		fmt.Println("The type is bool.")
	default:
		fmt.Println("Unknown type.")
	}
}

// 與 GenericTypes 做對照版本
func Int32AndFloat32InAssert(nums []interface{}) interface{} {
	var sum float32
	for _, num := range nums {
		// 需要 type assertion 才能進行操作
		switch n := num.(type) {
		case int32:
			sum += float32(n)
		case float32:
			sum += n
		}
	}

	// 如果傳入元素是int32 -> return []int32
	if len(nums) > 0 {
		if _, ok := nums[0].(int32); ok {
			return int32(sum)
		}
	}
	return sum
}

func Assertion() {

	/* Test: getType */
	var item interface{} = 10
	getType(item)

	/* Test: Int32AndFloat32InAssert */
	numsInt := []int32{1, 2, 3, 4, 5}
	numsFloat := []float32{0.3, 1.8, 3.3, 4.6, 5.0}

	// 轉換至空介面 interface{}
	interfaceInt := make([]interface{}, len(numsInt))
	for idx := range numsInt {
		interfaceInt[idx] = numsInt[idx]
	}

	interfaceFloat := make([]interface{}, len(numsFloat))
	for idx := range numsFloat {
		interfaceFloat[idx] = numsFloat[idx]
	}

	sumInt := Int32AndFloat32InAssert(interfaceInt)
	sumFloat := Int32AndFloat32InAssert(interfaceFloat)

	// fmt 套件的 Printf() 會使用 reflect 套件轉型 -> sum回傳值依然是空介面 interface{}
	fmt.Printf("sumInt: %v (type: %T)\n", sumInt, sumInt)
	fmt.Printf("sumFloat: %v (type: %T)\n", sumFloat, sumFloat)
}
