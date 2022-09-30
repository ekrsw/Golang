package foo

import "testing"

var Debug bool = false

func TestReturnMin(t *testing.T) {
	if Debug {
		t.Skip("foo.ReturnMin skip")
	}

	b := ReturnMin()
	if b != 1 {
		t.Errorf("%v != %v", b, 1)
	}
}