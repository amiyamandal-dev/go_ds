package unionandintersectionofarray

import "fmt"

func Intersection(nums1 []int, nums2 []int) []int {
	result_map := map[int]bool{}
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for _, i := range nums1 {
		_, ok := result_map[i]
		if !ok {
			result_map[i] = false
		}
	}

	for _, i := range nums2 {
		v, ok := result_map[i]
		if ok {
			fmt.Println(v)
			result_map[i] = true
		}
	}
	v := make([]int, 0)
	for val, state := range result_map {
		if state {
			v = append(v, val)
		}
	}
	return v
}
