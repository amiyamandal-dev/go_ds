package kthsmallelement

func partition(a []int, low int, high int) int {
	pivot := a[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if a[j] <= pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

func QuickSort(a []int, low int, high int) {
	if low < high {
		p := partition(a, low, high)
		QuickSort(a, low, p-1)
		QuickSort(a, p+1, high)
	}
}

func QuickSortWrap(a []int) {
	high := len(a) - 1
	low := 0
	QuickSort(a, low, high)
}
