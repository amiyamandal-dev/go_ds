package palindromenumber

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	string_val := strconv.Itoa(x)
	mid := (len(string_val) - 1) / 2
	left := 0
	right := len(string_val) - 1
	for left <= mid {
		if string_val[left] != string_val[right] {
			return false
		}

		left++
		right--
	}

	return true
}
