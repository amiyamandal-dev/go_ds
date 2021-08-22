package kadensalgo

import "testing"

func TestMaxSubArray(t *testing.T) {
	a := []int{1, 2, 3, -2, 5}
	rez := MaxSubArray(a)
	if rez != 9 {
		t.Errorf("error in code")
	}
	a = []int{-2, -3, -4, -1}
	rez = MaxSubArray(a)
	if rez != -1 {
		t.Errorf("error in code")
	}

}
