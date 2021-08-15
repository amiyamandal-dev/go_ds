package minmax

func MinMax(a []int, i int, j int) (int, int) {
	if i == j {
		min := a[i]
		max := a[i]
		return min, max
	} else {
		var min, max int
		if i == j-1 {
			if a[i] > a[j] {
				max = a[i]
				min = a[j]
			} else {
				max = a[j]
				min = a[i]
			}
			return min, max
		} else {
			mid := (i + j) / 2
			min_1, max_1 := MinMax(a, i, mid)
			min_2, max_2 := MinMax(a, mid+1, j)
			if max_1 > max_2 {
				max = max_1
			} else {
				max = max_2
			}

			if min_1 > min_2 {
				min = min_2
			} else {
				min = min_1
			}

			return min, max
		}
	}
}
