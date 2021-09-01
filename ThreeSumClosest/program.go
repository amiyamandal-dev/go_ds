package threesumclosest

import "math"

func sumArray(a []int) int {
	sum_val := 0
	for _, i := range a {
		sum_val += i
	}
	return sum_val
}

func ThreeSumClosest(nums []int, target int) int {
	// wong answer
	var divided [][]int
	for i := 0; i < len(nums); i += 1 {
		end := i + 3
		if end > len(nums) {
			end = len(nums)
		}
		if len(nums[i:end]) == 3 {
			divided = append(divided, nums[i:end])
		}
	}
	min_sum := sumArray(divided[0])
	for _, arr := range divided {
		temp_sum := sumArray(arr)
		if math.Abs(float64(target-min_sum)) > math.Abs(float64(target-temp_sum)) {
			min_sum = temp_sum
		}
	}
	return min_sum
}
