package movenegetivenumber

func MoveNegetiveNumber(a []int) {
	high := len(a) - 1
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			for j := high; j > i; j-- {
				if a[i] > a[j] {
					a[i], a[j] = a[j], a[i]
				}
			}
		}
	}
}
