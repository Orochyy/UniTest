package sample

import "testing"

func TestAdd(t *testing.T) {
	var x, y, z int = 1, 2, 3

	if z != Add(x, y) {
		t.Errorf("%d != %d", z, x+y)
	}
}
