package kthsmallelement

import "testing"

func TestSort(t *testing.T) {
	a := [...]int{9, 8, 1, 2, 7, 6, 3, 4, 5}
	QuickSortWrap(a[:])
	rez := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if a != rez {
		t.Errorf("error in code")
	}
}

func TestKthElement(t *testing.T) {
	a := [...]int{7, 10, 4, 3, 20, 15}
	K := 3
	rez := KthSmallElement(a[:], K)
	if rez != 7 {
		t.Errorf("error in code")
	}

}
