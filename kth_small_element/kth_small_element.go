package kthsmallelement

func KthSmallElement(a []int, kth int) int {
	QuickSortWrap(a)
	return a[kth-1]
}
