package minmax

import "testing"

func TestMinMax(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	min, max := MinMax(a[:], 0, 8)
	println(min)
	println(max)
	if min != 1 || max != 9 {
		t.Errorf("error in code")
	}

}
