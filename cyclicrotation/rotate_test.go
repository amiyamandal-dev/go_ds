package cyclicrotation

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {

	a := []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(a[:], 3)
	rez := []int{5, 6, 7, 1, 2, 3, 4}
	if !reflect.DeepEqual(rez, a) {
		t.Errorf("error in code")
	}
	b := []int{-1, -100, 3, 99}
	Rotate(b[:], 2)
	rez = []int{3, 99, -1, -100}
	if !reflect.DeepEqual(rez, b) {
		t.Errorf("error in code")
	}
}
