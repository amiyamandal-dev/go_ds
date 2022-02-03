package graphds

import "math"

func driver_code(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val.(int) <= min || root.Val.(int) >= max {
		return false
	}
	return driver_code(root.Left, min, root.Val.(int)) && driver_code(root.Right, root.Val.(int), max)
}

func IsValidBST(root *TreeNode) bool {
	return driver_code(root, math.MinInt, math.MaxInt)

}
