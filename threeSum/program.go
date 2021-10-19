package threesum

import (
	"sort"
)

func twoSumSomeExtraStep(nums []int, start int, end int, target_val int, to_be_added int) [][]int {
	// O(logn)
	result := make([][]int, 0)
	old_start := -999999
	old_end := -999999

	for start < end {
		if start != end {
			temp_sum := nums[start] + nums[end] + to_be_added
			if temp_sum == 0 {
				if old_end != nums[end] || old_start != nums[start] {
					result = append(result, []int{nums[start], nums[end], to_be_added})
				}
				old_start = nums[start]
				old_end = nums[end]
				start++
				end--
			} else if temp_sum > target_val {
				end--
			} else if temp_sum < target_val {
				start++
			}
		}

	}
	return result

}

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	} else if len(nums) == 3 {
		sum := nums[0] + nums[1] + nums[2]
		if sum == 0 {
			result := make([][]int, 0)
			result = append(result, []int{nums[0], nums[1], nums[2]})
			return result
		} else {
			return [][]int{}
		}
	}
	sort.Ints(nums) // O(nlogn)
	result := make([][]int, 0)
	len_of_array := len(nums) - 1
	old_val := nums[0]
	for i := 0; i < (len_of_array - 2); i++ {
		if i == 0 {
			old_val = nums[i]
			start := i + 1
			end := len_of_array
			k := nums[i]
			rez := twoSumSomeExtraStep(nums, start, end, 0, k)
			if len(rez) != 0 {
				result = append(result, rez...)
			}

		} else {
			if old_val != nums[i] {
				old_val = nums[i]
				start := i + 1
				end := len_of_array
				k := nums[i]
				rez := twoSumSomeExtraStep(nums, start, end, 0, k)
				if len(rez) != 0 {
					result = append(result, rez...)
				}
			}
		}

	}
	return result
}
