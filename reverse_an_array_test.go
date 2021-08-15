package main

import "testing"

func TestReverseAnArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ReverseAnArray(a[:])
	if a != [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1} {
		t.Errorf("error in code")
	}
}
