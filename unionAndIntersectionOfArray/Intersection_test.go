package unionandintersectionofarray

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	a := [...]int{1, 2, 2, 1}
	b := [...]int{2, 2}
	c := Intersection(a[:], b[:])
	rez := []int{2}
	if !reflect.DeepEqual(c, rez) {
		t.Errorf("error in code")
	}
}
