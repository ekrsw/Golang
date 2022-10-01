package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	if Debug {
		t.Skip("スキップします。")
	}

	i := 1

	v := IsOne(i)
	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}
