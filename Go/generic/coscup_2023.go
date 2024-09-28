package generic

import (
	"fmt"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// go: constraints
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}

// Need to implement all of them (SampleTypeA, SampleTypeB and Concat)
type SampleType[T any] interface {
	SampleTypeA | SampleTypeB
	Concat(T) T
}

type SampleTypeA struct {
	Str string
}

func (st SampleTypeA) Concat(s string) string {
	return st.Str + s
}

type SampleTypeB struct {
	Int int
}

func (st SampleTypeB) Concat(i int) int {
	return st.Int + i
}

func ConcatSampleType[T ~string | ~int, ST SampleType[T]](st ST, v T) T {
	return st.Concat(v)
}

type List[T constraints.Ordered] struct {
	list []T
}

func (ls *List[T]) Append(v T) {
	ls.list = append(ls.list, v)
}
func (ls *List[T]) All() []T {
	return ls.list
}

func (ls *List[T]) String() string {
	return fmt.Sprintf("%v", ls.list)
}

func ToString[T constraints.Ordered, R ~string | ~[]byte](ls List[T]) R {
	return R(fmt.Sprintf("%v", ls.list))
}

/** Why generic in go so slow ?

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type itab struct {
	inter *interfacetype // offset 0
	_type *_type         // offset 8
	hash  uint32         // offset 16
	_     [4]byte
	fun   [1]uintptr	 // offset 24
}
*/

/* Excpetion: string and []byte */
type ByteSeq interface {
	~string | ~[]byte
}

func SequenceToString[T ByteSeq](x T) string {
	ix := any(x)
	switch ix.(type) {
	case string:
		return ix.(string)
	case []byte:
		p := ix.([]byte)
		return *(*string)(unsafe.Pointer(&p))
	default:
		return ""
	}
}

/* Exception: callback */

// Callbacks: inlined function -> good efficiency
func MapInt(a []int, callback func(int) int) []int {
	for n, elem := range a {
		a[n] = callback(elem)
	}
	return a
}

// Callbacks: parametrize callback parameter -> less efficiency
func MapAnyA[I any](a []I, callback func(I) I) []I {
	for n, elem := range a {
		a[n] = callback(elem)
	}
	return a
}

// Callbacks: parametrize callback and its parameter -> good efficiency
func MapAnyB[I any, F func(I) I](a []I, callback F) []I {
	for n, elem := range a {
		a[n] = callback(elem)
	}
	return a
}

/* Conclusion

- Collaction
	T cannot be pointer or interface
- Callback
	Do parametrize callback type
- Serializers
	Performance down
- Datastores
	Performance down
	I/O latency >>>>>> code performance
- Mocks
	Pratically impossible
*/
