package threesum

import "testing"

func TestThreeSum(t *testing.T) {
	// a := [...]int{-1, 0, 1, 2, -1, -4}
	// rez := ThreeSum(a[:])
	// println(rez)
	a := []int{-2, 0, 0, 2, 2}
	rez := ThreeSum(a[:])
	println(rez)

}
