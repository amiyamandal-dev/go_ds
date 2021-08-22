package cyclicrotation

func Rotate(nums []int, k int) {
	len_val := len(nums)
	temp := make([]int, len_val)
	copy(temp, nums)
	for i := 0; i < len_val; i++ {
		nums[(k+i)%len_val] = temp[i]
	}
}
