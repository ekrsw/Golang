package alib

import "testing"

var Debug bool = false

func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("Average is skipped")
	}

	sl := []int{1, 2, 3, 4, 5}
	v := Average(sl)
	if 3 != v {
		t.Errorf("%v != %v", v, 3)
	}
}