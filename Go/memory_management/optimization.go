package memorymanagement

import (
	"strings"
	"sync"
)

// (a) Reduce the generation of temporary variables

func InefficientConcat(a, b string) string {
	return a + b
}

func EfficientConcat(a, b string) string {
	var buf strings.Builder
	buf.WriteString(a)
	buf.WriteString(b)
	return buf.String()
}

// (b) sync.Pool

var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024)
	},
}

func ProcessData() {
	buf := bufferPool.Get().(*[]byte)
	defer bufferPool.Put(buf)
}

// (c) Structure fragmentation
type InefficientStruct struct {
	FrequentField1 int
	FrequentField2 int
	RareField1     string
	RareField2     string
}

type FrequentFields struct {
	Field1 int
	Field2 int
}

type RareFields struct {
	Field1 string
	Field2 string
}

type EfficientStruct struct {
	Frequent FrequentFields
	Rare     *RareFields // Load the rare fields only when needed
}
