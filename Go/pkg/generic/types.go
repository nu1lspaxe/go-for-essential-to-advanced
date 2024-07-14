/*
	Go 泛型教學文章 : https://alankrantas.medium.com/%E7%B0%A1%E5%96%AE%E7%8E%A9-go-1-18-%E6%B3%9B%E5%9E%8B-1d09da07b70

	在 Go 語言裡，任何型別只要實作一個介面要求的所有方法，就可被視為符合該介面型別（隱性實作）。

	「interface{}」(空介面)，沒有定義任何方法，任何型別都可代入空介面。
	一旦值被放入空介面，Go 就無法得知值得 dynamic type (動態型別)，若想再知道只有兩種方式 :
		1. Type Assertion (型別斷言)
		2. reflect 套件

	泛型型別本身是介面，但其內容若是只有型別就可省略 interface{...}
	[T int] = [T interface{int}]
	[T ~int] = [T interface{~int}]	// ~int 表示底層型別為 int 的型別都算在內
	[T int | string] = [T interface{int | string}]
	[T any] = [T interface{}] 	// T 為任何型別
	[T comparable]	// T 為可比較的型別，若是符合以下條件也算:
											1. T 不是介面且支援 == / != / || (不是 >/<)
											2. T 是介面且底下所有型別都符合 comparable

*/

package generic

import (
	"fmt"
)

/* 聯集也可寫成介面 -> int32 | float32 */
type Int32AndFloat32 interface {
	int32 | float32
}

// string 型別也支援 "+" 運算子 -> [T int32 | float32 | string]
// bool 型別不支援 -> [T int32 | float32 | string | bool] -> 會報錯
func Int32AndFloat32InGeneric[T Int32AndFloat32](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}

func AssertAndGeneric[T int32 | float32](nums []T) T {
	var sum float32
	for _, num := range nums {
		switch n := any(num).(type) {
		case int32:
			sum += float32(n)
		case float32:
			sum += n
		}
	}
	return T(sum)
}

func GenericTypes() {
	numsInt := []int32{1, 2, 3, 4, 5}
	numsFloat := []float32{0.3, 1.8, 3.3, 4.6, 5.0}

	sumInt := Int32AndFloat32InGeneric[int32](numsInt)
	sumFloat := Int32AndFloat32InGeneric[float32](numsFloat)

	fmt.Printf("sumInt: %v (type: %T)\n", sumInt, sumInt)
	fmt.Printf("sumFloat: %v (type: %T)\n", sumFloat, sumFloat)
}
