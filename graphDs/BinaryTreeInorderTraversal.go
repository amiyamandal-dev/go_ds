package graphds

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	if root.Left != nil {
		res = inorderTraversal(root.Left)
	}
	res = append(res, root.Val.(int))
	if root.Right != nil {
		temp := inorderTraversal(root.Right)
		res = append(res, temp...)
	}
	return res
}
