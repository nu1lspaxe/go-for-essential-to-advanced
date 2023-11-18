package advanced

import "testing"

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

func foo(a, b int) bool {
	return a == b
}

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
