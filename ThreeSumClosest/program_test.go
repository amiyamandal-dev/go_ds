package threesumclosest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	a := []int{1,1,1,1}
	rez := ThreeSumClosest(a[:],-100)
	if rez != 2 {
		t.Errorf("error in code")
	}
}

