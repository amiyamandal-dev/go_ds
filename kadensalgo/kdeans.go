package kadensalgo

func MaxSubArray(nums []int) int {
	max_sum := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		if max_sum < sum {
			max_sum = sum
		}
	}
	return max_sum
}
