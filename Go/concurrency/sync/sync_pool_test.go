package sync

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

var buf, _ = json.Marshal(Student{Name: "nu1lspaxe", Age: 20})

func Unmarshal() {
	stu := &Student{}
	json.Unmarshal(buf, stu)
}

// BenchmarkUnmershal-20｜18798｜71251 ns/op｜1400 B/op｜7 allocs/op
func BenchmarkUnmershal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

// BenchmarkUnmarshalWithPool-20｜17310｜88698 ns/op｜248 B/op｜6 allocs/op
func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

// BenchmarkBufferWithPool-20｜21505221｜72.55 ns/op｜0 B/op｜0 allocs/op
func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

// BenchmarkBuffer-20｜645674｜1589 ns/op｜10240 B/op｜1 allocs/op
func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}
