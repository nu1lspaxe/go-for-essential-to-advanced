package advanced

import (
	"errors"
	"testing"
	"unicode/utf8"
)

/*
	Tutorial: Comprehensive Guide to Testing in Go (https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go)
*/

/*
	檔名須以 {filename}_test.go 命名。

	func Test(t *testing.T) {
		t.Log("Similar to fmt.Println() and concurrently safe")
		t.Fail() -> will show a test case has failed in the result
		t.FailNow() -> t.Fail() + safely exit without continuing
		t.Error("t.Log() + t.Fail()")
		t.Fatal("t.Log() + t.FailNow()")
	}
*/

/*
	Use build-in `Go: Toggle Test Coverage In Current Package` by `ctrl + shift + p` to search it
	Or `go test -coverprofile="c.out"` and `go tool cover -html="c.out"` to open and look file on browser
*/

func foo(a, b int) bool {
	return a == b
}

func reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

/* regular test(*testing.T), run command `go test .` */
func TestFoo(t *testing.T) {
	t.Run("Equal Case", func(t *testing.T) {
		if !foo(5, 5) {
			t.Error("they are not equal.")
		}
	})
	t.Run("Not Equal Case", func(t *testing.T) {
		if !foo(2, 3) {
			t.Log("they are not equal.")
		}
	})
}

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev, err := reverse(tc.in)
		if err != nil {
			return
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

/* benchamrk test(*testing.B), run command `go test -bench [file]` */
func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo(i, i)
	}
}

/* fuzzing test(*testing.F), run command `go test -fuzz=Fuzz [file] -fuzztime=[time]` */
/* fuzzing test generates test cases, especially for developer to test edge case. */

/* More fuzzing test tutorial:
- 善用 Go Fuzzing，幫助你寫出更完整的單元測試 (https://medium.com/starbugs/utilize-go-fuzzing-to-write-better-unit-tests-80bd37cd4e38)
- Tutorial: Getting started with fuzzing (https://go.dev/doc/tutorial/fuzz)
*/

func FuzzReverse(f *testing.F) {
	testcases := []string{"你好", "Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := reverse(orig)
		if err != nil {
			return
		}
		doubleRev, err := reverse(rev)
		if err != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
