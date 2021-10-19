package threesumclosest

import (
	"math"
	"sort"
)

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func twoSumSomeExtraStep(nums []int, start int, end int, target_val int, to_be_added int) int {
	// O(n)L
	result_val := nums[start] + nums[end] + to_be_added
	old_start := -999999
	old_end := -999999
	for start < end {
		if start != end {

			temp_sum := nums[start] + nums[end] + to_be_added
			if temp_sum > target_val {
				end--
			} else if temp_sum < target_val {
				start++
			}
			diff_1 := math.Abs(float64(temp_sum) - float64(target_val))
			diff_2 := math.Abs(float64(result_val) - float64(target_val))
			if diff_2 > diff_1 {
				if old_end != nums[end] || old_start != nums[start] {
					result_val = int(diff_1)
					old_start = nums[start]
					old_end = nums[end]
				}

			}
		}

	}
	return result_val

}

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	} else if len(nums) == 3 {
		sum := nums[0] + nums[1] + nums[2]
		return sum
	}
	sort.Ints(nums) // O(nlogn)
	len_of_array := len(nums) - 1
	all_rez := make([]int, 0)
	for i := 0; i < (len_of_array - 2); i++ {
		start := i + 1
		end := len_of_array
		k := nums[i]
		rez := twoSumSomeExtraStep(nums, start, end, target, k)
		all_rez = append(all_rez, rez)
	}
	min_val, _ := MinMax(all_rez)
	return min_val
}
